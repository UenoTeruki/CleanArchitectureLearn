package gateway

import (
	"clean-architecture-learn/entity"
	"clean-architecture-learn/usecase/port"
	"log"

	"github.com/labstack/echo/v4"
)

type TodoRepository struct {
	ctx echo.Context
}

func NewTodoRepository(ctx echo.Context) port.RepositoryPort {
	return &TodoRepository{
		ctx: ctx,
	}
}

func(r *TodoRepository) GetAllTodos() []*entity.ToDo {
	return []*entity.ToDo{
		{ Title: "title1", Content: "content1" },
		{ Title: "title2", Content: "content2" },
		{ Title: "title3", Content: "content3" },
	}
}

func(r *TodoRepository) CreateNewTodo() *entity.ToDo {
	var todo entity.ToDo

	err := r.ctx.Bind(&todo)
	if err != nil {
		log.Fatal(err)
	}

	return &todo
}
