package driver

import (
	"clean-architecture-learn/adapter/controller"
	"clean-architecture-learn/adapter/gateway"
	"clean-architecture-learn/adapter/presenter"
	"clean-architecture-learn/usecase/interactor"

	"github.com/labstack/echo/v4"
)

func Routing(e *echo.Echo) {
	api := e.Group("/api")

	toDoController := controller.ToDoController{
		InputFactory: interactor.NewTodoInteractor,
		OutputFactory: presenter.NewTodoPresenter,
		RepositoryFactory: gateway.NewTodoRepository,
	}
	api.GET("/todos", toDoController.GetAllTodos)
}
