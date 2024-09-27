package middlewares

import (
	log "log"
	time "time"

	gin "github.com/gin-gonic/gin"
)

// LoggingMiddleware logs every request with details like method, URL, and status code
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Allow the request to proceed to the next middleware or handler
		c.Next()

		// After handler finishes, log the details
		log.Printf("Request: %s %s | Status: %d | Duration: %s", c.Request.Method, c.Request.URL, c.Writer.Status(), time.Since(startTime))
	}
}
