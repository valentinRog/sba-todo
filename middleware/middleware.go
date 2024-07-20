package middleware

import (
	"context"

	"github.com/valentinRog/sba-todo/handler"
	"github.com/valentinRog/sba-todo/middleware/auth"
	"github.com/valentinRog/sba-todo/store"
)

type Middleware struct {
	Auth auth.Middleware
}

func New(ctx context.Context, h *handler.Handlers, q *store.Queries) *Middleware {
	return &Middleware{
		Auth: *auth.New(ctx, h, q),
	}
}
