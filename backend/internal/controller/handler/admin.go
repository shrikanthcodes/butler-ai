package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Admin handlers for admin routes
func BanUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "User has been banned"})
}

func GetAllConversations(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "All conversations"})
}

func GetAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "All users"})
}
