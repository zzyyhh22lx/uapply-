package cors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Cors Set up cross-domain
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// Release all OPTIONS methods
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// handle the requests
		c.Next()
	}
}
