package service

import (
	"todo"
	"todo/pkg/repository"
)

// Называю интерфейсы исходя из их доменной зоны

type Authorization interface {
	CreateUser(user todo.User) (int, error) //возвращает айди созданного пользователя
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error) // Извлекаю id из payload
}

type TodoList interface{}

type TodoItem interface{}

// Собирает все сервисы в одном месте

type Service struct {
	Authorization
	TodoList
	TodoItem
}

// Конструктор принимает указатель на бд, это и есть внедрение зависимостей

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
