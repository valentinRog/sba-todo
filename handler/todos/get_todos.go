package todos

import (
	"github.com/labstack/echo/v4"
	"github.com/valentinRog/sba-todo/ui/templates/layout"
	"github.com/valentinRog/sba-todo/ui/templates/todos"
)

func GetTodos(c echo.Context) error {
	return layout.Layout(todos.Todos()).Render(c.Response())
}
