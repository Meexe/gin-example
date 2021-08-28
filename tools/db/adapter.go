package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const connStr = "user=maxpak dbname=hack sslmode=disable"

func New() (*sql.DB, error) {
	return sql.Open("postgres", connStr)
}
