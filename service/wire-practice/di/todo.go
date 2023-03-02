package di

import (
	"github.com/google/wire"
	"go-playground/service/wire-practice/handler"
	"go-playground/service/wire-practice/infra"
	"go-playground/service/wire-practice/repository"
	"go-playground/service/wire-practice/usercase"
)

var todoProviderSet = wire.NewSet(
	infra.HttpProviderSet,
	infra.MongoProviderSet,
	repository.HttpProviderSet,
	repository.MongoProviderSet,
	usercase.TodoProviderSet,
	handler.TodoProviderSet,
)
