package repository

import (
	"github.com/google/wire"
	"go-playground/service/wire-practice/infra"
)

type HttpRepositoryInterface interface {
	Get() string
}

type HttpRepository struct {
	client *infra.HttpClient
}

func NewHttpRepository(client *infra.HttpClient) *HttpRepository {
	return &HttpRepository{
		client: client,
	}
}

func (h *HttpRepository) Get() string {
	return "123123"
}

var HttpProviderSet = wire.NewSet(NewHttpRepository)
