package handler

import (
	"context"

	"github.com/valentinRog/sba-todo/handler/auth"
	"github.com/valentinRog/sba-todo/handler/login"
	"github.com/valentinRog/sba-todo/handler/static"
	"github.com/valentinRog/sba-todo/handler/todos"
	"github.com/valentinRog/sba-todo/store"
)

type Handlers struct {
	Login  login.Handlers
	Todos  todos.Handlers
	Static static.Handlers
	Auth   auth.Handlers
}

func New(ctx context.Context, q *store.Queries) *Handlers {
	return &Handlers{
		Login: *login.New(ctx, q),
		Todos: *todos.New(ctx, q),
		Auth:  *auth.New(ctx, q),
	}
}
