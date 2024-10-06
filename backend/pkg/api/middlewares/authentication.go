package middlewares

import (
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

// General Authentication Middleware (for users)
func GeneralAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// token := c.Request.Header.Get("Authorization")

		// // Mock token validation (in a real case, you'd validate JWT or session token)
		// if token != "Bearer valid-token" {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		// 	c.Abort() // Stop further processing
		// 	return
		// }
		// // Continue to the next middleware or the handler
		c.Next()
	}
}

// Admin Authentication Middleware
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		// Mock token validation (in a real case, you'd validate JWT or session token)
		if token != "Bearer valid-admin-token" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort() // Stop further processing
			return
		}
		// Continue to the next middleware or the handler
		c.Next()
	}
}
