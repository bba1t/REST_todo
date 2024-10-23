package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"todo"
	"todo/pkg/repository"
)

const (
	salt       = "asassd22234h3j"   // Набор символов для усложнения хеширования
	signingKey = "sfdfe345IOHUIYY3" // Ключ подписи, используется для расшифровки токена
)

type tokenClaims struct { // дополнение к обычным claims
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

/*
Поле с репозиторием нужно для реализации архитектуры, сервис должен обращаться к бд.
По логике должно быть так: repo repository.Repository, но нет смысла обращаться ко всей структуре,
когда мне нужно только одно ее поле, его и укажу: repo repository.Authorization
*/

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// Просто отправляю данные о пользователе дальше вниз, в репозиторий

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	// если такой пользователь существует, генерирую токен
	// первый аргумент - метод для подписи(шифрования), второй - данные токена
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(), // токен проживет 12 часов с момента создания
			IssuedAt:  time.Now().Unix(),                     // время создания токена
		},
		user.Id,
	})
	return token.SignedString([]byte(signingKey))
}

// Извлекаю id из payload
func (s *AuthService) ParseToken(accessToken string) (int, error) {
	// 2. Проверка валидности, парсинг

	// jwt.ParseWithClaims разбирает токен accessToken в &tokenClaims{}, который хранит claims (например, UserId)
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// проверяет, что метод подписи является HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		// если метод подписи правильный, возвращаю секретный ключ подписи, он используется для проверки
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err // если токен, невалиден или подпись кривая
	}

	// извлекаются claims из токена и приводятся к типу *tokenClaims
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	// возвращаю id пользователя
	return claims.UserId, nil
}

// Функция хеширования пароля
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
