package todos

import (
	"github.com/labstack/echo/v4"
	"github.com/valentinRog/sba-todo/store/todo"
	"github.com/valentinRog/sba-todo/ui/templates/todos"
)

func PostTodo(c echo.Context) error {
	todo.AddTodo(c.FormValue("text"))
	return todos.TodoList().Render(c.Response())
}
