package store

import (
	"context"
	"database/sql"
	"log"

	_ "embed"

	"github.com/valentinRog/sba-todo/store/todo"
	"github.com/valentinRog/sba-todo/store/user"
	_ "modernc.org/sqlite"
)

type Queries struct {
	User *user.Queries
	Todo *todo.Queries
}

func NewQueries(db *sql.DB) *Queries {
	return &Queries{
		User: user.New(db),
		Todo: todo.New(db),
	}
}

func Init(ctx context.Context) *sql.DB {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	user.Init(ctx, db)
	todo.Init(ctx, db)

	return db
}
