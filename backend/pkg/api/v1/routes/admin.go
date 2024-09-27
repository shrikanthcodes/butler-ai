package routes

import (
	middleware "backend/pkg/api/middlewares"
	handler "backend/pkg/api/v1/handlers"

	gin "github.com/gin-gonic/gin"
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
