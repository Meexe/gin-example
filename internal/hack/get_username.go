package hack

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserName returns username
func (s *Service) GetUsername(c *gin.Context) {
	user := c.Params.ByName("name")
	c.JSON(http.StatusOK, gin.H{"user": user})
}
