package repository

import "go-playground/service/wire-practice/infra"

type HttpRepository struct {
	client *infra.HttpClient
}

func NewHttpRepository(client *infra.HttpClient) *HttpRepository {
	return &HttpRepository{
		client: client,
	}
}

func (h *HttpRepository) Get() string {
	return "123"
}
