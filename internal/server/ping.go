package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping test
func (s *Server) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
