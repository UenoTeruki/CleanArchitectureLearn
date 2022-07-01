package controller

import (
	"clean-architecture-learn/usecase/port"

	"github.com/labstack/echo/v4"
)

type TodoController struct {
	InputFactory func(presenter port.TodoOutputPort, repository port.TodoRepository) port.TodoInputPort
	OutputFactory func(ctx echo.Context) port.TodoOutputPort
	RepositoryFactory func(ctx echo.Context) port.TodoRepository
}

func(c *TodoController) GetAllTodos(ctx echo.Context) error {
	presenter := c.OutputFactory(ctx)
	repository := c.RepositoryFactory(ctx)

	interactor := c.InputFactory(presenter, repository)
	return interactor.GetAllTodos()
}

func(c *TodoController) CreateNewTodo(ctx echo.Context) error {
	presenter := c.OutputFactory(ctx)
	repository := c.RepositoryFactory(ctx)

	interactor := c.InputFactory(presenter, repository)
	return interactor.CreateNewTodo()
}
