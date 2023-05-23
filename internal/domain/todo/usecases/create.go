package usecases

import (
	"github.com/rudolfoborges/google-wire-sample/internal/domain/todo/entities"
	"github.com/rudolfoborges/google-wire-sample/internal/domain/todo/repositories"
)

type CreateTodo struct {
	TodoRepository repositories.TodoRepository
}

type CreateTodoResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func NewCreateTodoUseCase(todoRepository repositories.TodoRepository) *CreateTodo {
	return &CreateTodo{
		TodoRepository: todoRepository,
	}
}

func (u *CreateTodo) Execute(title, description string) {
	todo := entities.NewTodo(title, description)
	u.TodoRepository.Create(todo)
}
