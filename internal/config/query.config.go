package config

import (
	"database/sql"

	queries "github.com/hyper-dot/books-server/internal/generated"
)

func Query() (*queries.Queries, *sql.DB) {
	db := ConnectDatabase()
	queries := queries.New(db)
	return queries, db
}
