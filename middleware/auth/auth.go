package auth

import (
	"context"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/valentinRog/sba-todo/handler"
	"github.com/valentinRog/sba-todo/store"
	"github.com/valentinRog/sba-todo/ui/templates/layout"
	logintmpl "github.com/valentinRog/sba-todo/ui/templates/login"
)

type Middleware struct {
	q   *store.Queries
	h   *handler.Handlers
	ctx context.Context
}

func New(ctx context.Context, h *handler.Handlers, q *store.Queries) *Middleware {
	return &Middleware{
		ctx: ctx,
		h:   h,
		q:   q,
	}
}

func (m *Middleware) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, _ := c.Request().Cookie("token")
		fmt.Println(c.Path())
		if strings.HasPrefix(c.Path(), "/static") {
			return next(c)
		}
		if cookie == nil && !strings.HasPrefix(c.Path(), "/login") && !strings.HasPrefix(c.Path(), "/auth") {
			c.Response().Header().Set("HX-Push-Url", "/login")
			if c.Request().Header.Get("HX-Request") == "true" {
				return logintmpl.LoginPage().Render(c.Response())
			}
			return layout.Layout(logintmpl.LoginPage()).Render(c.Response())
		}
		if cookie != nil && strings.HasPrefix(c.Path(), "/login") {
			c.Response().Header().Set("HX-Push-Url", "/")
			return m.h.Todos.GetTodos(c)
		}
		return next(c)
	}
}
