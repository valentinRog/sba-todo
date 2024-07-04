package login

import (
	"github.com/labstack/echo/v4"
	"github.com/valentinRog/sba-todo/ui/templates/layout"
	"github.com/valentinRog/sba-todo/ui/templates/login"
)

func GetLogin(c echo.Context) error {
	if c.Request().Header.Get("HX-Request") == "true" {
		return login.LoginPage().Render(c.Response())
	}
	return layout.Layout(login.LoginPage()).Render(c.Response())
}
