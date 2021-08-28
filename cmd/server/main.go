package main

import (
	"log"

	"github.com/Meexe/gin-example/internal/hack"
	"github.com/Meexe/gin-example/tools"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	db, err := tools.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	srv := hack.New(db)

	server := initServer(srv)
	server.Run(":8000")
}

func initServer(s *hack.Service) *gin.Engine {
	a := gin.Default()

	a.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	a.GET("/ping", s.Ping)
	a.GET("/user/:name", s.GetUsername)
	a.GET("/org-structure", s.GetOrgStructure)

	return a
}
