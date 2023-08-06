package main

import (
	"todolist/handlers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/todos", handlers.GetTodos)
	r.POST("/todos/add", handlers.AddTodo)
	// wip: Add routes for updating, deleting later

	r.Run(":8080")
}
