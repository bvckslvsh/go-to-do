package service

import (
	gotodo "github.com/bvckslvsh/go-to-do"
	"github.com/bvckslvsh/go-to-do/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list gotodo.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]gotodo.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int) (gotodo.TodoList, error) {
	return s.repo.GetById(userId, listId)
}
