package todo

import (
	"github.com/google/wire"
	"github.com/rudolfoborges/google-wire-sample/internal/dataproviders"
	"github.com/rudolfoborges/google-wire-sample/internal/domain/todo/controllers"
	"github.com/rudolfoborges/google-wire-sample/internal/domain/todo/usecases"
)

// Container is the container for the todo domain
var TodoSuperSet = wire.NewSet(
	dataproviders.NewTodoRepository,
	usecases.NewFetchTodosUseCase,
	usecases.NewCompleteTodoUseCase,
	usecases.NewCreateTodoUseCase,
	controllers.NewTodoController,
)
