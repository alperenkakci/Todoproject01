package todo

import (
	"net/http"

	"todoproject01/middleware"
	"todoproject01/service/todo"

	"github.com/gin-gonic/gin"
)

type TodoService interface {
	GetTodoList(id string) (todo.Todo, bool)
	CreateTodoList(id string) (todo.Todo, error)
	UpdateTodoList(id string, todo todo.Todo) (todo.Todo, error)
	DeleteTodoList(id string) error
	AddTodoItem(listID string, item todo.TodoMessage) (todo.Todo, error)
}

type Controller struct {
	service TodoService
}

type TodoParams struct {
	ID string `uri:"id" binding:"required"`
}

func NewController(service TodoService) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) CreateTodoList(cx *gin.Context) {
	roles := []string{middleware.UserRole, middleware.AdminRole}
	if err := middleware.AuthzMiddleware(roles...)(cx); err != nil {
		cx.JSON(http.StatusUnauthorized, gin.H{"error": "Insufficient permissions"})
		return
	}

	var params TodoParams
	if err := cx.BindJSON(&params); err != nil {
		cx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	t, err := c.service.CreateTodoList(params.ID)
	if err != nil {
		cx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	cx.JSON(http.StatusOK, t)
}

func (c *Controller) GetTodoList(cx *gin.Context) {
	roles := []string{middleware.UserRole, middleware.AdminRole}
	if err := middleware.AuthzMiddleware(roles...)(cx); err != nil {
		cx.JSON(http.StatusUnauthorized, gin.H{"error": "Insufficient permissions"})
		return
	}

	var params TodoParams
	if err := cx.ShouldBindUri(&params); err != nil {
		cx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	td, exist := c.service.GetTodoList(params.ID)
	if !exist {
		cx.JSON(http.StatusNotFound, map[string]string{"message": "This todo list not found"})
		return
	}

	cx.JSON(http.StatusOK, td)
}

func (c *Controller) UpdateTodoList(cx *gin.Context) {
	var params TodoParams
	if err := cx.ShouldBindUri(&params); err != nil {
		cx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	username, exists := cx.Get("username")
	if !exists || username != params.ID {
		cx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var todo todo.Todo
	if err := cx.BindJSON(&todo); err != nil {
		cx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	updatedTodo, err := c.service.UpdateTodoList(params.ID, todo)
	if err != nil {
		cx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	cx.JSON(http.StatusOK, updatedTodo)
}

func (c *Controller) DeleteTodoList(cx *gin.Context) {
	var params TodoParams
	if err := cx.ShouldBindUri(&params); err != nil {
		cx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	username, exists := cx.Get("username")
	if !exists || username != params.ID {
		cx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err := c.service.DeleteTodoList(params.ID)
	if err != nil {
		cx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	cx.JSON(http.StatusOK, map[string]string{"message": "Todo list deleted"})
}

func (c *Controller) AddTodoItem(cx *gin.Context) {
	var params TodoParams
	if err := cx.ShouldBindUri(&params); err != nil {
		cx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	username, exists := cx.Get("username")
	if !exists || username != params.ID {
		cx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var item todo.TodoMessage
	if err := cx.BindJSON(&item); err != nil {
		cx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	item.TODOListID = params.ID
	todoList, err := c.service.AddTodoItem(params.ID, item)
	if err != nil {
		cx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	cx.JSON(http.StatusOK, todoList)
}
