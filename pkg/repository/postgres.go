package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

// Тут будут параметры необходимые для подключения к бд
type Config struct {
	Username string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
	Password string
}

// Реализую подключение к бд
func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	// Открываю бд с нужными данными
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	// Проверяю подключение
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
