package interactor

import (
	"clean-architecture-learn/usecase/port"
)

type ToDoInteractor struct {
	presenter port.TodoOutputPort
	repository port.TodoRepository
}

func NewTodoInteractor(presenter port.TodoOutputPort, repository port.TodoRepository) port.TodoInputPort {
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
