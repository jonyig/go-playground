package repository

import "go-playground/service/wire-practice/infra"

type TodoRepository struct {
	client *infra.MongoClient
}

func NewTodoRepository(client *infra.MongoClient) *TodoRepository {
	return &TodoRepository{
		client: client,
	}
}

func (t *TodoRepository) Get() string {
	return "789"
}
