package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {
	connectionStr := "postgres://roshanpaudel:@localhost:5432/mydb?sslmode=disable"
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatal("Database connection error :", err)
	}

	return db
}
