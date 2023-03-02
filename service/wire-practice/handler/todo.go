package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go-playground/service/wire-practice/usercase"
)

type TodoHandler struct {
	TodoUserCase *usercase.TodoUserCase
}

func NewTodoRoute(h *TodoHandler, r *gin.Engine) {
	r.GET("/todo", h.Get)
}

func NewTodoHandler(r *gin.Engine, userCase *usercase.TodoUserCase) *TodoHandler {
	h := &TodoHandler{
		TodoUserCase: userCase,
	}

	NewTodoRoute(h, r)

	return h
}

func (t *TodoHandler) Get(c *gin.Context) {
	c.JSON(200, t.TodoUserCase.Get())
	return
}

var TodoProviderSet = wire.NewSet(NewTodoHandler)
