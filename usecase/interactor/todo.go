package interactor

import (
	"clean-architecture-learn/usecase/port"
)

type ToDoInteractor struct {
	presenter port.OutputPort
	repository port.RepositoryPort
}

func NewTodoInteractor(presenter port.OutputPort, repository port.RepositoryPort) port.InputPort {
	return &ToDoInteractor{
		presenter: presenter,
		repository: repository,
	}
}

func(i *ToDoInteractor) GetAllTodos() error {
	todos := i.repository.GetAllTodos()
	return i.presenter.OutputAllTodos(todos)
}

func(i *ToDoInteractor) CreateNewTodo() error {
	todo := i.repository.CreateNewTodo()
	return i.presenter.SuccessfullyCreated(todo)
}
