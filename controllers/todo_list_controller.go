package controllers

import (
	"todo-list/services"

	"github.com/gin-gonic/gin"
)

func GetAllTodosController(c *gin.Context) {
	todolist := services.GetAllTodos()
	c.JSON(200, todolist)
}
