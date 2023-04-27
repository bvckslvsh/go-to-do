package repository

import (
	gotodo "github.com/bvckslvsh/go-to-do"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user gotodo.User) (int, error)
	GetUser(username, password string) (gotodo.User, error)
}

type TodoList interface {
	Create(userId int, list gotodo.TodoList) (int, error)
	GetAll(userId int) ([]gotodo.TodoList, error)
	GetById(userId, listId int) (gotodo.TodoList, error)
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoItem
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
