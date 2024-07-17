package store

import (
	"context"
	"database/sql"
	"log"

	_ "embed"
	_ "modernc.org/sqlite"
)

//go:embed todo/schema.sql
var ddl string

func Init(ctx context.Context) *sql.DB {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := db.ExecContext(ctx, ddl); err != nil {
		log.Fatal(err)
	}

	return db
}
