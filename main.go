package main

import (
	"todo-list/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.TodoListRoutes(r)
	r.Run(":8080")
}
