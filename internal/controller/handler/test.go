package handler

import (
	"github.com/gin-gonic/gin"
)

// HealthCheck returns the list of users from the db
func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "API is up and running, rejoice!",
	})
}
