package service

import (
	gotodo "github.com/bvckslvsh/go-to-do"
	"github.com/bvckslvsh/go-to-do/pkg/repository"
)

type Authorization interface {
	CreateUser(user gotodo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
