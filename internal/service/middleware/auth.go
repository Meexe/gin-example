package middleware

import "github.com/gin-gonic/gin"

// Authorized group (uses gin.BasicAuth() middleware)
// Same than:
// authorized := r.Group("/")
// authorized.Use(gin.BasicAuth(gin.Credentials{
//	  "foo":  "bar",
//	  "manu": "123",
//}))

var Auth = gin.BasicAuth(gin.Accounts{
	"foo":  "bar", // user:foo password:bar
	"manu": "123", // user:manu password:123
})
