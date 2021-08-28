package hack

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping test
func (s *Service) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
