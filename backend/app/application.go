package app

import (
	"todolist/app/repositories"
	"todolist/app/handlers"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

type Application struct {
	Router *gin.Engine
}

func NewApplication(db *sql.DB) *Application {

	todoRepository := repositories.NewTodoRepository(db)
	todoHandler := handlers.NewTodoHandler(*todoRepository)

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowCredentials = true
	router.Use(cors.New(config))
	
	router.GET("/todos", todoHandler.GetTodos)
	router.POST("/todos/add", todoHandler.AddTodo)
	router.DELETE("/todos/:id", todoHandler.DeleteTodo)
	router.PUT("/todos/:id", todoHandler.UpdateTodo)
	
	return &Application{
		Router: router,
	}
}

func (app *Application) Run(addr string) error {
	return app.Router.Run(addr)
}
