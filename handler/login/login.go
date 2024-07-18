package login

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/valentinRog/sba-todo/store"
	"github.com/valentinRog/sba-todo/ui/templates"
)

var tmpl templates.Templates

type Handlers struct {
	q   *store.Queries
	ctx context.Context
}

func New(ctx context.Context, q *store.Queries) *Handlers {
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
	listTodos, _ := h.q.Todo.ListTodos(h.ctx)
	return tmpl.Todos.Todos(listTodos).Render(c.Response())
}

func (h *Handlers) GetSigninForm(c echo.Context) error {
	return tmpl.Login.SigninForm().Render(c.Response())
}

func (h *Handlers) GetSignupForm(c echo.Context) error {
	return tmpl.Login.SignupForm().Render(c.Response())
}
