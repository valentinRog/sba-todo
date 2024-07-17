package handler

import (
	"context"

	"github.com/valentinRog/sba-todo/handler/login"
	"github.com/valentinRog/sba-todo/handler/todos"
	"github.com/valentinRog/sba-todo/store"
)

type Handlers struct {
	Login login.Handlers
	Todos todos.Handlers
}

func New(ctx context.Context, q *store.Queries) *Handlers {
	return &Handlers{
		Login: *login.New(ctx, q.User),
		Todos: *todos.New(ctx, q.Todo),
	}
}
