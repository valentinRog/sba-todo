package todos

import (
	"github.com/labstack/echo/v4"
	"github.com/valentinRog/spa-todo/ui/templates/layout"
)

func GetTodos(c echo.Context) error {
	return layout.Layout().Render(c.Response())
}
