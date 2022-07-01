package port

import "clean-architecture-learn/entity"

type TodoInputPort interface {
	GetAllTodos() error
	CreateNewTodo() error
}

type TodoOutputPort interface {
	OutputAllTodos(todos []*entity.Todo) error
	SuccessfullyCreated(todo *entity.Todo) error
}

type TodoRepository interface {
	GetAllTodos() []*entity.Todo
	CreateNewTodo() *entity.Todo
}
