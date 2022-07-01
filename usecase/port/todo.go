package port

import "clean-architecture-learn/entity"

type InputPort interface {
	GetAllTodos() error
	CreateNewTodo() error
}

type OutputPort interface {
	OutputAllTodos(todos []*entity.ToDo) error
	SuccessfullyCreated(todo *entity.ToDo) error
}

type RepositoryPort interface {
	GetAllTodos() []*entity.ToDo
	CreateNewTodo() *entity.ToDo
}
