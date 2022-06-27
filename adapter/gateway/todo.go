package gateway

import (
	"clean-architecture-learn/entity"
	"clean-architecture-learn/usecase/port"
)

type TodoRepository struct {}

func NewTodoRepository() port.RepositoryPort {
	return &TodoRepository{}
}

func(r *TodoRepository) GetAllTodos() []*entity.ToDo {
	return []*entity.ToDo{
		{ Title: "title1", Content: "content1" },
		{ Title: "title2", Content: "content2" },
		{ Title: "title3", Content: "content3" },
	}
}
