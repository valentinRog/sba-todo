package todos

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/valentinRog/sba-todo/store/todo"
	"github.com/valentinRog/sba-todo/ui/templates/layout"
	"github.com/valentinRog/sba-todo/ui/templates/todos"
)

func GetTodos(c echo.Context) error {
	if c.Request().Header.Get("HX-Request") == "true" {
		return todos.Todos().Render(c.Response())
	}
	return layout.Layout(todos.Todos()).Render(c.Response())
}

func PostAddTodo(c echo.Context) error {
	todo.AddTodo(c.FormValue("text"))
	return todos.TodoList().Render(c.Response())
}

func PostDeleteTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	todo.DeleteTodo(id)
	return todos.TodoList().Render(c.Response())
}
