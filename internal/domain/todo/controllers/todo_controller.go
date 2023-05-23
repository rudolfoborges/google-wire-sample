package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rudolfoborges/google-wire-sample/internal/domain/todo/usecases"
)

type TodoController struct {
	CreateTodo   *usecases.CreateTodo
	FetchTodos   *usecases.FetchTodos
	CompleteTodo *usecases.CompleteTodo
}

type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewTodoController(
	createTodo *usecases.CreateTodo,
	fetchTodos *usecases.FetchTodos,
	completeTodo *usecases.CompleteTodo,
) *TodoController {
	return &TodoController{
		CreateTodo:   createTodo,
		FetchTodos:   fetchTodos,
		CompleteTodo: completeTodo,
	}
}

func (c *TodoController) DoFetchTodos(ctx *fiber.Ctx) error {
	todos := c.FetchTodos.Execute()
	return ctx.JSON(todos)
}

func (c *TodoController) DoCreateTodo(ctx *fiber.Ctx) error {
	var request CreateTodoRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	c.CreateTodo.Execute(request.Title, request.Description)

	return ctx.SendStatus(fiber.StatusCreated)
}

func (c *TodoController) DoCompleteTodo(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")

	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}

	c.CompleteTodo.Execute(id)

	return ctx.SendStatus(fiber.StatusOK)
}
