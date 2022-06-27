package port

import "clean-architecture-learn/entity"

type InputPort interface {
	GetAllTodos() error
}

type OutputPort interface {
	OutputAllTodos(todos []*entity.ToDo) error
}

type RepositoryPort interface {
	GetAllTodos() []*entity.ToDo
}
