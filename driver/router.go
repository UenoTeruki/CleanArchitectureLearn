package driver

import (
	"clean-architecture-learn/driver/router"

	"github.com/labstack/echo/v4"
)

type Router interface {
	Routing(g *echo.Group)
}

func setRouting(r Router, g *echo.Group) {
	r.Routing(g)
}

func Routing(e *echo.Echo) {
	g := e.Group("/api")

	setRouting(router.TodoRouter{}, g)
}
