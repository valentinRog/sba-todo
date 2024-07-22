package todos

import (
	"context"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/valentinRog/sba-todo/auth"
	"github.com/valentinRog/sba-todo/store"
	todostore "github.com/valentinRog/sba-todo/store/todo"
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
	token, _ := c.Request().Cookie("token")
	userID, _ := auth.GetUserIdFromToken(token.Value)
	user, _ := h.q.User.GetUser(h.ctx, userID)
	listTodos, _ := h.q.Todo.ListTodos(h.ctx, userID)
	return layout.Layout(
		homelayout.Layout(user.Username, todotmpl.Todos(listTodos)),
	).Render(c.Response())
}

func (h *Handlers) PostAddTodo(c echo.Context) error {
	token, _ := c.Request().Cookie("token")
	userID, _ := auth.GetUserIdFromToken(token.Value)
	_, err := h.q.Todo.CreateTodo(h.ctx, todostore.CreateTodoParams{
		UserID: userID,
		Name:   c.FormValue("text"),
	})
	if err != nil {
		log.Fatal(err)
	}
	listTodos, _ := h.q.Todo.ListTodos(h.ctx, userID)
	return todotmpl.TodoList(listTodos).Render(c.Response())
}

func (h *Handlers) PostDeleteTodo(c echo.Context) error {
	token, _ := c.Request().Cookie("token")
	userID, _ := auth.GetUserIdFromToken(token.Value)
	id, _ := strconv.Atoi(c.Param("id"))
	h.q.Todo.DeleteTodo(h.ctx, int64(id))
	listTodos, _ := h.q.Todo.ListTodos(h.ctx, userID)
	return todotmpl.TodoList(listTodos).Render(c.Response())
}
