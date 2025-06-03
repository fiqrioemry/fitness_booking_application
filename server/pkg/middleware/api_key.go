package middleware

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func APIKeyGateway(skippedPaths []string) gin.HandlerFunc {
	requiredKey := os.Getenv("API_KEY")

	return func(c *gin.Context) {
		for _, path := range skippedPaths {
			if strings.HasPrefix(c.Request.URL.Path, path) {
				c.Next()
				return
			}
		}

		apiKey := c.GetHeader("X-API-KEY")
		if apiKey == "" || apiKey != requiredKey {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized - invalid or missing API key"})
			return
		}

		c.Next()
	}
}
