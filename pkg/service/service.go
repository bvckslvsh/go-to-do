package service

import (
	gotodo "github.com/bvckslvsh/go-to-do"
	"github.com/bvckslvsh/go-to-do/pkg/repository"
)

type Authorization interface {
	CreateUser(user gotodo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list gotodo.TodoList) (int, error)
	GetAll(userId int) ([]gotodo.TodoList, error)
	GetById(userId, listId int) (gotodo.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input gotodo.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item gotodo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]gotodo.TodoItem, error)
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
