package handlers

import (
	"strconv"
	"net/http"
	"todolist/models"
	"github.com/gin-gonic/gin"
)

var todos = []models.Todo{
	{ID: 1, Task: "Buy milk", Status: false},
	{ID: 2, Task: "Buy eggs", Status: false},
} // In-memory store for now for simplicity

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

func DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	for index, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:index], todos[index+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	var updatedTodo models.Todo
	if err := c.BindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for index, todo := range todos {
		if todo.ID == id {
			todos[index] = updatedTodo
			c.JSON(http.StatusOK, updatedTodo)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}
