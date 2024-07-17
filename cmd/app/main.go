package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/valentinRog/sba-todo/handler"
	"github.com/valentinRog/sba-todo/handler/static"
	"github.com/valentinRog/sba-todo/store"
)

func main() {
	ctx := context.Background()

	e := echo.New()
	db := store.Init(ctx)
	queries := store.NewQueries(db)
	handlers := handler.New(ctx, queries)

	e.GET("/", handlers.Todos.GetTodos)
	e.POST("/add-todo", handlers.Todos.PostAddTodo)
	e.POST("/delete-todo/:id", handlers.Todos.PostDeleteTodo)
	e.GET("/style", static.GetStyle)
	e.GET("/htmx", static.GetHtmx)
	e.GET("/login", handlers.Login.GetLogin)
	e.POST("/signup", handlers.Login.PostSignup)

	e.Logger.Fatal(e.Start(":80"))
}
