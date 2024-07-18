package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/valentinRog/sba-todo/handler"
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
	e.GET("/style", handlers.Static.GetStyle)
	e.GET("/htmx", handlers.Static.GetHtmx)
	e.GET("/login", handlers.Login.GetLogin)
	e.GET("/login/signin-form", handlers.Login.GetSigninForm)
	e.GET("/login/signup-form", handlers.Login.GetSignupForm)
	e.POST("/signup", handlers.Login.PostSignup)

	e.Logger.Fatal(e.Start(":80"))
}
