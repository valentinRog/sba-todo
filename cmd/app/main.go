package main

import (
	"github.com/labstack/echo/v4"
	"github.com/valentinRog/sba-todo/handler/static"
	"github.com/valentinRog/sba-todo/handler/todos"
)

func main() {
	e := echo.New()

	e.GET("/", todos.GetTodos)
	e.GET("/style", static.HandleStyle)
	e.GET("/htmx", static.HandleHtmx)

	e.Logger.Fatal(e.Start(":80"))
}
