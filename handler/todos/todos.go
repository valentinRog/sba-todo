package todos

import (
	"context"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/valentinRog/sba-todo/auth"
	"github.com/valentinRog/sba-todo/store"
	homelayout "github.com/valentinRog/sba-todo/ui/templates/homelayout"
	"github.com/valentinRog/sba-todo/ui/templates/layout"
	todotmpl "github.com/valentinRog/sba-todo/ui/templates/todos"
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

func (h *Handlers) GetTodos(c echo.Context) error {
	listTodos, _ := h.q.Todo.ListTodos(h.ctx)
	token, _ := c.Request().Cookie("token")
	user, _ := h.q.User.GetUser(h.ctx, auth.GetUserIdFromToken(token.Value))
	return layout.Layout(
		homelayout.Layout(user.Username, todotmpl.Todos(listTodos)),
	).Render(c.Response())
}

func (h *Handlers) PostAddTodo(c echo.Context) error {
	_, err := h.q.Todo.CreateTodo(h.ctx, c.FormValue("text"))
	if err != nil {
		log.Fatal(err)
	}
	listTodos, _ := h.q.Todo.ListTodos(h.ctx)
	return todotmpl.TodoList(listTodos).Render(c.Response())
}

func (h *Handlers) PostDeleteTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	h.q.Todo.DeleteTodo(h.ctx, int64(id))
	listTodos, _ := h.q.Todo.ListTodos(h.ctx)
	return todotmpl.TodoList(listTodos).Render(c.Response())
}
