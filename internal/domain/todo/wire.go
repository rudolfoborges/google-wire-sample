//go:build wireinject
// +build wireinject

package todo

import (
	"github.com/google/wire"
	"github.com/rudolfoborges/google-wire-sample/internal/domain/todo/controllers"
)

func NewTodoController() *controllers.TodoController {
	wire.Build(TodoSuperSet)
	return &controllers.TodoController{}
}
