package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func New() (*sql.DB, error) {
	return sql.Open("postgres", os.Getenv("DATABASE_URL"))
}
