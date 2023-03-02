package repository

import (
	"github.com/google/wire"
	"go-playground/service/wire-practice/infra"
)

type TodoRepository struct {
	client *infra.MongoClient
}

func NewTodoRepository(client *infra.MongoClient) *TodoRepository {
	return &TodoRepository{
		client: client,
	}
}

func (t *TodoRepository) Get() string {
	return "78913"
}

var MongoProviderSet = wire.NewSet(NewTodoRepository)
