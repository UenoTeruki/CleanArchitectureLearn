package router

import(
	"clean-architecture-learn/adapter/controller"
	"clean-architecture-learn/adapter/gateway"
	"clean-architecture-learn/adapter/presenter"
	"clean-architecture-learn/usecase/interactor"

	"github.com/labstack/echo/v4"
)

type TodoRouter struct {}

func(r TodoRouter) Routing(g *echo.Group) {
	toDoController := controller.TodoController{
		InputFactory: interactor.NewTodoInteractor,
		OutputFactory: presenter.NewTodoPresenter,
		RepositoryFactory: gateway.NewTodoRepository,
	}

	g.GET("/todos", toDoController.GetAllTodos)
	g.POST("/todos", toDoController.CreateNewTodo)
}
