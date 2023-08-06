package main

import (
	"todolist/handlers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	r.GET("/todos", handlers.GetTodos)
	r.POST("/todos/add", handlers.AddTodo)
	r.DELETE("/todos/:id", handlers.DeleteTodo)
	r.PUT("/todos/:id", handlers.UpdateTodo)

	r.Run(":8080")
}
