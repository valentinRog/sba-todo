package login

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/valentinRog/sba-todo/store"
	"github.com/valentinRog/sba-todo/ui/templates/layout"
	logintmpl "github.com/valentinRog/sba-todo/ui/templates/login"
)

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
		return logintmpl.LoginPage().Render(c.Response())
	}
	return layout.Layout(logintmpl.LoginPage()).Render(c.Response())
}

func (h *Handlers) GetSigninForm(c echo.Context) error {
	return logintmpl.SigninForm().Render(c.Response())
}

func (h *Handlers) GetSignupForm(c echo.Context) error {
	return logintmpl.SignupForm().Render(c.Response())
}
