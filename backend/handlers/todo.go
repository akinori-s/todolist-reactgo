package handlers

import (
	"net/http"
	"todolist/models"
	"github.com/gin-gonic/gin"
)

var todos = []models.Todo{} // In-memory store for now for simplicity

func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func AddTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.BindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todos = append(todos, todo)
	c.JSON(http.StatusOK, todo)
}

// wip: adding handlers like UpdateTodo, DeleteTodo later
