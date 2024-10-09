package v1

import (
	gin "github.com/gin-gonic/gin"
	"github.com/shrikanthcodes/butler-ai/backend/internal/controller/handler"
	middleware "github.com/shrikanthcodes/butler-ai/backend/internal/controller/http"
)

func AdminRoutes(router *gin.Engine) {
	admin := router.Group("/api/admin")
	admin.Use(middleware.AdminAuthMiddleware())
	{
		admin.PUT("/user/{id}/ban", handler.BanUser)             // Ban a user
		admin.GET("/conversations", handler.GetAllConversations) // Get all conversations
		admin.GET("/users", handler.GetAllUsers)                 // Get all users
	}
}
