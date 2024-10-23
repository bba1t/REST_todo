package repository

import (
	"github.com/jmoiron/sqlx"
	"todo"
)

// Называю интерфейсы исходя из их доменной зоны

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error) // todo.User и есть токен
}

type TodoList interface{}

type TodoItem interface{}

// Собирает все сервисы в одном месте

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// Конструктор

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
