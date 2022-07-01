package presenter

import (
	"clean-architecture-learn/entity"
	"clean-architecture-learn/usecase/port"

	"net/http"
	"github.com/labstack/echo/v4"
)

type TodoPresenter struct {
	ctx echo.Context
}

func NewTodoPresenter(ctx echo.Context) port.TodoOutputPort {
	return &TodoPresenter{
		ctx: ctx,
	}
}

func(p *TodoPresenter) OutputAllTodos(todos []*entity.Todo) error {
	return p.ctx.JSON(http.StatusOK, todos)
}

func(p *TodoPresenter) SuccessfullyCreated(todo *entity.Todo) error {
	return p.ctx.JSON(http.StatusOK, todo)
}