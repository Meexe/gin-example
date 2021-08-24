package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* example curl for /admin with basicauth header
   Zm9vOmJhcg== is base64("foo:bar")

	curl -X POST \
  	http://localhost:8080/admin \
  	-H 'authorization: Basic Zm9vOmJhcg==' \
  	-H 'content-type: application/json' \
  	-d '{"value":"bar"}'
*/

func IsAdmin(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)

	// Parse JSON
	var json struct {
		Value string `json:"value" binding:"required"`
	}

	if c.Bind(&json) == nil {
		db[user] = json.Value
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
