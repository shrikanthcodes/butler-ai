package middlewares

import (
	"math/rand"
	"time"

	config "backend/internal/config"
	models "backend/pkg/models"

	"github.com/gin-gonic/gin"
)

// HealthCheck returns the list of users from the database
func HealthCheck(c *gin.Context) {
	var users []models.User
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

// CreateRandomUser inserts a random user into the database
func CreateRandomUser(c *gin.Context) {
	// Create a random user
	user := models.User{
		Name:  randomString(5), // Generate a random 5-letter name
		Email: randomEmail(),   // Generate a random email
	}

	// Insert into database
	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Random user created successfully!",
		"user":    user,
	})
}

// randomString generates a random string of the given length
func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}

// randomEmail generates a random email address
func randomEmail() string {
	return randomString(5) + "@example.com"
}
