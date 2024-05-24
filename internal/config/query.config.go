package config

import (
	"database/sql"

	"github.com/hyper-dot/books-server/internal/queries"
)

func Query() (*queries.Queries, *sql.DB) {
	db := ConnectDatabase()
	queries := queries.New(db)
	return queries, db
}
