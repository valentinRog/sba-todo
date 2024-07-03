package main

import (
	"github.com/labstack/echo/v4"
	"github.com/valentinRog/spa-todo/handler/static"
	"github.com/valentinRog/spa-todo/handler/todos"
)

func main() {
	e := echo.New()

	e.GET("/", todos.GetTodos)
	e.GET("/style", static.HandleStyle)

	e.Logger.Fatal(e.Start(":80"))
}
