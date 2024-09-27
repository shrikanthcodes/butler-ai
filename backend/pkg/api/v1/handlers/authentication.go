package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Auth handlers for auth routes
func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Login"})
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Register"})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Logout"})
}

func RefreshToken(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Refresh token"})
}

func ResetPassword(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Reset password"})
}
