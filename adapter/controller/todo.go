package controller

import (
	"clean-architecture-learn/usecase/port"

	"github.com/labstack/echo/v4"
)

type ToDoController struct {
	InputFactory func(presenter port.OutputPort, repository port.RepositoryPort) port.InputPort
	OutputFactory func(ctx echo.Context) port.OutputPort
	RepositoryFactory func() port.RepositoryPort
}

func(c *ToDoController) GetAllTodos(ctx echo.Context) error {
	presenter := c.OutputFactory(ctx)
	repository := c.RepositoryFactory()

	interactor := c.InputFactory(presenter, repository)
	return interactor.GetAllTodos()
}
