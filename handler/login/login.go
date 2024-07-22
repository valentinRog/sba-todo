package login

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/valentinRog/sba-todo/store"
	"github.com/valentinRog/sba-todo/ui/templates/layout"
	logintmpl "github.com/valentinRog/sba-todo/ui/templates/login"
	signintmpl "github.com/valentinRog/sba-todo/ui/templates/login/signin"
	signuptmpl "github.com/valentinRog/sba-todo/ui/templates/login/signup"
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

func (h *Handlers) GetSigninForm(c echo.Context) error {
	return signintmpl.SigninForm().Render(c.Response())
}

func (h *Handlers) GetSignupForm(c echo.Context) error {
	return signuptmpl.SignupForm().Render(c.Response())
}

func (h *Handlers) GetLogin(c echo.Context) error {
	return layout.Layout(logintmpl.LoginPage(signintmpl.SigninForm())).
		Render(c.Response())
}
