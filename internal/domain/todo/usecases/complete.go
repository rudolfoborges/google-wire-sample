package usecases

import "github.com/rudolfoborges/google-wire-sample/internal/domain/todo/repositories"

type CompleteTodo struct {
	TodoRepository repositories.TodoRepository
}

func NewCompleteTodoUseCase(todoRepository repositories.TodoRepository) *CompleteTodo {
	return &CompleteTodo{
		TodoRepository: todoRepository,
	}
}

func (u *CompleteTodo) Execute(id int) {
	todo := u.TodoRepository.Find(id)
	todo.Complete()
	u.TodoRepository.Update(todo)
}
