package main

import (
	todoController "todoproject01/controller/todo"
	"todoproject01/middleware"
	todoRepository "todoproject01/repository/todo"
	todoService "todoproject01/service/todo"

	"github.com/gin-gonic/gin"
)

func main() {
	todoMap := make(map[string]todoService.Todo)

	todoRepo := todoRepository.NewRepository(todoMap)
	todoSvc := todoService.NewService(todoRepo)
	todoCntrl := todoController.NewController(todoSvc)

	router := gin.Default()

	router.POST("/login", middleware.LoginHandler)

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/todos", todoCntrl.CreateTodoList)
		protected.GET("/todos/:id", todoCntrl.GetTodoList)
		// Add other CRUD endpoints here
	}

	router.Run("localhost:8080")
}
