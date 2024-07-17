package login

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/valentinRog/sba-todo/store/user"
	"github.com/valentinRog/sba-todo/ui/templates/layout"
	"github.com/valentinRog/sba-todo/ui/templates/login"
)

type Handlers struct {
	q   *user.Queries
	ctx context.Context
}

func New(ctx context.Context, q *user.Queries) *Handlers {
	return &Handlers{
		ctx: ctx,
		q:   q,
	}
}

func (h *Handlers) GetLogin(c echo.Context) error {
	if c.Request().Header.Get("HX-Request") == "true" {
		return login.LoginPage().Render(c.Response())
	}
	return layout.Layout(login.LoginPage()).Render(c.Response())
}

func (h *Handlers) PostSignup(c echo.Context) error {
	// return layout.Layout(todos.Todos(listTodos))
	return layout.Layout(login.LoginPage()).Render(c.Response())
}
