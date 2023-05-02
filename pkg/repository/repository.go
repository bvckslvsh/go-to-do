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
	Delete(userId, listId int) error
	Update(userId, listId int, input gotodo.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item gotodo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]gotodo.TodoItem, error)
	GetById(userId, itemId int) (gotodo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input gotodo.UpdateItemInput) error
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
		TodoItem:      NewTodoItemPostgres(db),
	}
}
