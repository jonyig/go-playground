package handler

import (
	"github.com/gin-gonic/gin"
	"go-playground/service/wire-practice/usercase"
)

type TodoHandler struct {
	TodoUserCase *usercase.TodoUserCase
}

func NewTodoHandler(userCase *usercase.TodoUserCase) *TodoHandler {
	return &TodoHandler{
		TodoUserCase: userCase,
	}
}

func (t *TodoHandler) Get(c *gin.Context) {
	c.JSON(200, t.TodoUserCase.Get())
	return
}
