package usercase

import (
	"fmt"
	"github.com/google/wire"
	"go-playground/service/wire-practice/repository"
)

type TodoUserCaseInterface interface {
	Get() string
}

type TodoUserCase struct {
	TodoRepository repository.TodoRepositoryInterface
	HttpRepository repository.HttpRepositoryInterface
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

var TodoProviderSet = wire.NewSet(NewTodoUserCase)
