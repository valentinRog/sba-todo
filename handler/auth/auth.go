package auth

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/valentinRog/sba-todo/auth"
	"github.com/valentinRog/sba-todo/store"
	userstore "github.com/valentinRog/sba-todo/store/user"
	todotmpl "github.com/valentinRog/sba-todo/ui/templates/todos"
	"net/http"
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

func (h *Handlers) PostSignup(c echo.Context) error {
	name := c.FormValue("username")
	password := c.FormValue("password")
	userId, _ := h.q.User.CreateUser(h.ctx, userstore.CreateUserParams{Username: name, Password: password})
	token := auth.CreateSession(userId)
	cookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	}
	c.SetCookie(&cookie)
	listTodos, _ := h.q.Todo.ListTodos(h.ctx)
	return todotmpl.Todos(listTodos).Render(c.Response())
}
