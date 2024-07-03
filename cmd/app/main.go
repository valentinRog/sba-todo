package main

import (
	"github.com/labstack/echo/v4"
	"github.com/valentinRog/sba-todo/handler/static"
	"github.com/valentinRog/sba-todo/handler/todos"
	"github.com/valentinRog/sba-todo/store/todo"
)

func main() {
	e := echo.New()

	todo.AddTodo("un truc")
	todo.AddTodo("un autre truc")

	e.GET("/", todos.GetTodos)
	e.POST("/addtodo", todos.PostTodo)
	e.GET("/style", static.HandleStyle)
	e.GET("/htmx", static.HandleHtmx)

	e.Logger.Fatal(e.Start(":80"))
}
