//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"go-playground/service/wire-practice/handler"
)

func CreateTodoHandler() *handler.TodoHandler {
	panic(wire.Build(todoProviderSet))
}
