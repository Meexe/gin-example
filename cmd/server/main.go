package main

import (
	"database/sql"
	"log"

	"github.com/Meexe/gin-example/internal/server"
	_ "github.com/lib/pq"
)

var connStr = "user=maxpak dbname=hack sslmode=disable"

func main() {

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	s := server.New(db)
	s.Run(":8000")
}
