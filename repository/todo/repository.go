package todo

import (
	"fmt"
	"time"
	"todoproject01/service/todo"
)

type Repository struct {
	todoMap map[string]todo.Todo
}

func NewRepository(todoMap map[string]todo.Todo) *Repository {
	return &Repository{
		todoMap: todoMap,
	}
}

func (s *Repository) GetTodoList(id string) (todo.Todo, bool) {
	td, exist := s.todoMap[id]
	if !exist || td.DeletionDate != nil {
		return todo.Todo{}, false
	}
	return td, true
}

func (s *Repository) CreateTodoList(todo todo.Todo) (todo.Todo, error) {
	s.todoMap[todo.ID] = todo
	return todo, nil
}

func (s *Repository) UpdateTodoList(todo todo.Todo) (todo.Todo, error) {
	s.todoMap[todo.ID] = todo
	return todo, nil
}

func (s *Repository) DeleteTodoList(id string) error {
	if todo, exists := s.todoMap[id]; exists {
		now := time.Now()
		todo.DeletionDate = &now
		s.todoMap[id] = todo
		return nil
	}
	return fmt.Errorf("todo list not found")
}
