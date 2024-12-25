package routes

import (
	"todo-list/controllers"

	"github.com/gin-gonic/gin"
)

func TodoListRoutes(router *gin.Engine) {
	todolistRouters := router.Group("/todo-list")

	todolistRouters.GET("/", controllers.GetAllTodosController)
}
