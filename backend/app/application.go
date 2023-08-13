package app

import (
	"todolist/app/repositories"
	"todolist/app/usecases"
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

	authRepository := repositories.NewAuthRepository(db)
	authUsecase := usecases.NewAuthUsecase(*authRepository)
	authHandler := handlers.NewAuthHandler(*authUsecase)

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowCredentials = true
	router.Use(cors.New(config))
	
	router.GET("/todos", todoHandler.GetTodos)
	router.POST("/todos/add", todoHandler.AddTodo)
	router.DELETE("/todos/:id", todoHandler.DeleteTodo)
	router.PUT("/todos/:id", todoHandler.UpdateTodo)
	router.POST("/login", authHandler.Login)
	router.POST("/signup", authHandler.Signup)
	router.POST("/checkLogin", authHandler.CheckLogin)
	router.POST("/signout", authHandler.Signout)
	
	return &Application{
		Router: router,
	}
}

func (app *Application) Run(addr string) error {
	return app.Router.Run(addr)
}
