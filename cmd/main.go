package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rudolfoborges/google-wire-sample/internal/domain/todo"
)

func main() {
	controller := todo.NewTodoController()

	if controller == nil {
		panic("controller is nil")
	}

	app := fiber.New()

	app.Get("/todos", controller.DoFetchTodos)
	app.Post("/todos", controller.DoCreateTodo)
	app.Put("/todos/:id/complete", controller.DoCompleteTodo)

	app.Listen(":3000")
}
