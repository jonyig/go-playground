package usercase

import (
	"fmt"
	"go-playground/service/wire-practice/repository"
)

type TodoUserCase struct {
	TodoRepository *repository.TodoRepository
	HttpRepository *repository.HttpRepository
}

func NewTodoUserCase(
	todoRepository *repository.TodoRepository,
	httpRepository *repository.HttpRepository,
) *TodoUserCase {
	return &TodoUserCase{
		TodoRepository: todoRepository,
		HttpRepository: httpRepository,
	}
}

func (t *TodoUserCase) Get() string {
	return fmt.Sprintf(
		"%s %s",
		t.TodoRepository.Get(),
		t.HttpRepository.Get(),
	)
}
