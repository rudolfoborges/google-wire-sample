package dataproviders

import (
	"github.com/rudolfoborges/google-wire-sample/internal/domain/todo/entities"
	"github.com/rudolfoborges/google-wire-sample/internal/domain/todo/repositories"
)

type TodoRepository struct{}

var todos []*entities.Todo

func NewTodoRepository() repositories.TodoRepository {
	return &TodoRepository{}
}

func (r *TodoRepository) Find(id int) *entities.Todo {
	return todos[id-1]
}

func (r *TodoRepository) FindAll() []*entities.Todo {
	return todos
}

func (r *TodoRepository) Create(todo *entities.Todo) {
	todo.ID = len(todos) + 1
	todos = append(todos, todo)
}

func (r *TodoRepository) Update(todo *entities.Todo) {
	todos[todo.ID-1] = todo
}
