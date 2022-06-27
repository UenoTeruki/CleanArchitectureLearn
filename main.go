package main

import(
	"clean-architecture-learn/driver"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  driver.Routing(e)

  e.Logger.Fatal(e.Start(":3000"))
}
