package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/valentinRog/sba-todo/handler"
	"github.com/valentinRog/sba-todo/middleware"
	"github.com/valentinRog/sba-todo/store"
)

func main() {
	ctx := context.Background()

	db := store.Init(ctx)
	queries := store.NewQueries(db)
	handlers := handler.New(ctx, queries)
	mw := middleware.New(ctx, handlers, queries)

	e := echo.New()
	e.Use(mw.Auth.AuthMiddleware)

	{
		g := e.Group("/login")
		g.GET("", handlers.Login.GetLogin)
		g.GET("/signin-form", handlers.Login.GetSigninForm)
		g.GET("/signup-form", handlers.Login.GetSignupForm)
	}
	{
		g := e.Group("/auth")
		g.POST("/signup", handlers.Auth.PostSignup)
		g.POST("/signin", handlers.Auth.PostSignin)
		g.POST("/logout", handlers.Auth.PostLogout)
	}
	{
		g := e.Group("/static")
		g.GET("/style", handlers.Static.GetCSS)
		g.GET("/htmx", handlers.Static.GetHtmx)
	}
	e.GET("/", handlers.Todos.GetTodos)
	e.POST("/add-todo", handlers.Todos.PostAddTodo)
	e.POST("/delete-todo/:id", handlers.Todos.PostDeleteTodo)

	e.Logger.Fatal(e.Start(":80"))
}
