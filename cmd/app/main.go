package main

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/valentinRog/sba-todo/handler"
	"github.com/valentinRog/sba-todo/store"
)

func mw(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("hello je suis le mw")
		return next(c)
	}
}

func main() {
	ctx := context.Background()

	db := store.Init(ctx)
	queries := store.NewQueries(db)
	handlers := handler.New(ctx, queries)

	e := echo.New()
	e.Use(mw)

	{
		g := e.Group("/login")
		g.GET("", handlers.Login.GetLogin)
		g.GET("/signin-form", handlers.Login.GetSigninForm)
		g.GET("/signup-form", handlers.Login.GetSignupForm)
	}

	e.GET("/", handlers.Todos.GetTodos)
	e.POST("/add-todo", handlers.Todos.PostAddTodo)
	e.POST("/delete-todo/:id", handlers.Todos.PostDeleteTodo)
	e.GET("/style", handlers.Static.GetStyle)
	e.GET("/htmx", handlers.Static.GetHtmx)
	e.POST("/signup", handlers.Auth.PostSignup)

	e.Logger.Fatal(e.Start(":80"))
}
