package main

import (
	"github.com/Meexe/gin-example/internal/server"
)

func main() {
	s := server.New()
	s.Run(":8080")
}
