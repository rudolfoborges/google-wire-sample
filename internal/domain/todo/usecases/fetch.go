package usecases

import "github.com/rudolfoborges/google-wire-sample/internal/domain/todo/repositories"

type FetchTodos struct {
	TodoRepository repositories.TodoRepository
}

type FetchTodosResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func NewFetchTodosUseCase(todoRepository repositories.TodoRepository) *FetchTodos {
	return &FetchTodos{
		TodoRepository: todoRepository,
	}
}

func (u *FetchTodos) Execute() []*FetchTodosResponse {
	todos := u.TodoRepository.FindAll()

	var response []*FetchTodosResponse

	for _, todo := range todos {
		response = append(response, &FetchTodosResponse{
			ID:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   todo.Completed,
		})
	}

	return response
}
