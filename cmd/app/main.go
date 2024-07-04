package main

import (
	"github.com/labstack/echo/v4"
	"github.com/valentinRog/sba-todo/handler/login"
	"github.com/valentinRog/sba-todo/handler/static"
	"github.com/valentinRog/sba-todo/handler/todos"
	"github.com/valentinRog/sba-todo/store/todo"
)

func main() {
	e := echo.New()

	todo.AddTodo("un truc")
	todo.AddTodo("un autre truc")

	e.GET("/", todos.GetTodos)
	e.POST("/add-todo", todos.PostAddTodo)
	e.POST("/delete-todo/:id", todos.PostDeleteTodo)
	e.GET("/style", static.GetStyle)
	e.GET("/htmx", static.GetHtmx)
	e.GET("/login", login.GetLogin)

	e.Logger.Fatal(e.Start(":80"))
}
