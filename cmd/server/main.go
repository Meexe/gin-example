package main

import (
	"log"

	"github.com/Meexe/gin-example/internal/hack"
	"github.com/Meexe/gin-example/tools/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	dbAdp, err := db.New()
	if err != nil {
		log.Fatal(err)
	}
	srv := hack.New(dbAdp)

	server := initServer(srv)
	server.Run(":8000")
}

func initServer(s *hack.Service) *gin.Engine {
	a := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	a.Use(cors.New(config))

	a.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	a.GET("/ping", s.Ping)

	a.GET("/org-structure", s.GetOrgStructure)
	a.GET("/search", s.Search)
	a.GET("/card/:ID", s.GetCard)
	a.POST("/task", s.SaveTask)

	return a
}
