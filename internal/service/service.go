package server

import (
	"github.com/Meexe/gin-example/internal/service/middleware"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func New() *gin.Engine {
	a := gin.Default()

	public := a.Group("/")
	authorized := a.Group("/admin", middleware.Auth)

	public.GET("ping", Ping)
	public.GET("user/:name", GetUsername)

	authorized.POST("me", IsAdmin)

	return a
}
