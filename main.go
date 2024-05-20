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
		protected.PUT("/todos/:id", todoCntrl.UpdateTodoList)
		protected.DELETE("/todos/:id", todoCntrl.DeleteTodoList)
		protected.POST("/todos/:id/items", todoCntrl.AddTodoItem)
	}

	router.Run("localhost:8080")
}
