package todo

import (
	"fmt"
	"time"
)

type TodoRepository interface {
	GetTodoList(id string) (Todo, bool)
	CreateTodoList(todo Todo) (Todo, error)
	UpdateTodoList(todo Todo) (Todo, error)
	DeleteTodoList(id string) error
}

type Todo struct {
	ID                string                 `json:"id"`
	CreationDate      time.Time              `json:"creationDate"`
	ModificationDate  time.Time              `json:"modificationDate"`
	DeletionDate      *time.Time             `json:"deletionDate"`
	CompletionPercent int                    `json:"completionPercent"`
	Messages          map[string]TodoMessage `json:"messages"`
}

type TodoMessage struct {
	ID               string     `json:"id"`
	TODOListID       string     `json:"todoListId"`
	CreationDate     time.Time  `json:"creationDate"`
	ModificationDate time.Time  `json:"modificationDate"`
	DeletionDate     *time.Time `json:"deletionDate"`
	Content          string     `json:"content"`
	CompletionStatus bool       `json:"completionStatus"`
}

type Service struct {
	repo TodoRepository
}

func NewService(repo TodoRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetTodoList(id string) (Todo, bool) {
	return s.repo.GetTodoList(id)
}

func (s *Service) CreateTodoList(id string) (Todo, error) {
	todo := Todo{
		ID:                id,
		CreationDate:      time.Now(),
		ModificationDate:  time.Now(),
		CompletionPercent: 0,
		Messages:          make(map[string]TodoMessage),
	}
	return s.repo.CreateTodoList(todo)
}

func (s *Service) UpdateTodoList(id string, todo Todo) (Todo, error) {
	todo.ModificationDate = time.Now()
	return s.repo.UpdateTodoList(todo)
}

func (s *Service) DeleteTodoList(id string) error {
	todo, exists := s.repo.GetTodoList(id)
	if !exists {
		return fmt.Errorf("todo list not found")
	}

	now := time.Now()
	todo.DeletionDate = &now
	_, err := s.repo.UpdateTodoList(todo)
	return err
}
