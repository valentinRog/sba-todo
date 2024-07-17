package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/valentinRog/sba-todo/handler/login"
	"github.com/valentinRog/sba-todo/handler/static"
	"github.com/valentinRog/sba-todo/handler/todos"
	"github.com/valentinRog/sba-todo/store"
	"github.com/valentinRog/sba-todo/store/todo"
)

func main() {
	ctx := context.Background()

	e := echo.New()

	todosHandlers := todos.New(ctx, todo.New(store.Init(ctx)))

	e.GET("/", todosHandlers.GetTodos)
	e.POST("/add-todo", todosHandlers.PostAddTodo)
	e.POST("/delete-todo/:id", todosHandlers.PostDeleteTodo)
	e.GET("/style", static.GetStyle)
	e.GET("/htmx", static.GetHtmx)
	e.GET("/login", login.GetLogin)

	e.Logger.Fatal(e.Start(":80"))
}
