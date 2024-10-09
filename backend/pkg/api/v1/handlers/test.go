package handlers

import (
	config "github.com/shrikanthcodes/butler-ai/backend/internal/config"
	model "github.com/shrikanthcodes/butler-ai/backend/pkg/models"

	"github.com/gin-gonic/gin"
)

// HealthCheck returns the list of users from the database
func HealthCheck(c *gin.Context) {
	var users []model.User
	result := config.DB.Find(&users) // Fetch all users from the DB

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "API is up and running",
		"users":  users,
	})
}
