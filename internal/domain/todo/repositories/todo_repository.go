package repositories

import "github.com/rudolfoborges/google-wire-sample/internal/domain/todo/entities"

type TodoRepository interface {
	Find(id int) *entities.Todo
	FindAll() []*entities.Todo
	Create(todo *entities.Todo)
	Update(todo *entities.Todo)
}
