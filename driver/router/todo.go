package router

import(
	"clean-architecture-learn/adapter/controller"
	"clean-architecture-learn/adapter/gateway"
	"clean-architecture-learn/adapter/presenter"
	"clean-architecture-learn/usecase/interactor"

	"github.com/labstack/echo/v4"
)

type ToDoRouter struct {}

func(r ToDoRouter) Routing(g *echo.Group) {
	toDoController := controller.ToDoController{
		InputFactory: interactor.NewTodoInteractor,
		OutputFactory: presenter.NewTodoPresenter,
		RepositoryFactory: gateway.NewTodoRepository,
	}

	g.GET("/todos", toDoController.GetAllTodos)
}
