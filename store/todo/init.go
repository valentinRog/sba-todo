package todo

import (
	"context"
	"database/sql"
	_ "embed"
	"log"
)

//go:embed schema.sql
var ddl string

func Init(ctx context.Context, db *sql.DB) {
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		log.Fatal(err)
	}
}
