package todos

import (
	"context"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/valentinRog/sba-todo/store/todo"
	"github.com/valentinRog/sba-todo/ui/templates/layout"
	"github.com/valentinRog/sba-todo/ui/templates/todos"
)

type Handlers struct {
	q   *todo.Queries
	ctx context.Context
}

func New(ctx context.Context, q *todo.Queries) *Handlers {
	return &Handlers{
		ctx: ctx,
		q:   q,
	}
}

func (h *Handlers) GetTodos(c echo.Context) error {
	listTodos, _ := h.q.ListTodos(h.ctx)
	if c.Request().Header.Get("HX-Request") == "true" {
		return todos.Todos(listTodos).Render(c.Response())
	}
	return layout.Layout(todos.Todos(listTodos)).Render(c.Response())
}

func (h *Handlers) PostAddTodo(c echo.Context) error {
	err := h.q.CreateTodo(h.ctx, c.FormValue("text"))
	if err != nil {
		log.Fatal(err)
	}
	listTodos, _ := h.q.ListTodos(h.ctx)
	return todos.TodoList(listTodos).Render(c.Response())
}

func (h *Handlers) PostDeleteTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	h.q.DeleteTodo(h.ctx, int64(id))
	listTodos, _ := h.q.ListTodos(h.ctx)
	return todos.TodoList(listTodos).Render(c.Response())
}
