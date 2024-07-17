package login

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/valentinRog/sba-todo/store/user"
	"github.com/valentinRog/sba-todo/ui/templates"
)

var tmpl templates.Templates

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
		return tmpl.Login.LoginPage().Render(c.Response())
	}
	return tmpl.Layout.Layout(tmpl.Login.LoginPage()).Render(c.Response())
}

func (h *Handlers) PostSignup(c echo.Context) error {
	// return layout.Layout(todos.Todos(listTodos))
	return tmpl.Layout.Layout(tmpl.Login.LoginPage()).Render(c.Response())
}
