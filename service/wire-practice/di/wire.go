//go:build wireinject
// +build wireinject

package di

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"go-playground/service/wire-practice/handler"
)

func InitTodoHandler(r *gin.Engine) *handler.TodoHandler {
	panic(wire.Build(todoProviderSet))
}
