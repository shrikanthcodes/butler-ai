package routes

import (
	middleware "github.com/shrikanthcodes/butler-ai/backend/pkg/api/middlewares"
	handler "github.com/shrikanthcodes/butler-ai/backend/pkg/api/v1/handlers"

	gin "github.com/gin-gonic/gin"
)

func ConversationRoutes(router *gin.Engine) {
	// Set the routes
	conv := router.Group("/api/conversation")
	conv.Use(middleware.GeneralAuthMiddleware())
	{
		conv.GET("/:id", handler.GetConversationByID)                  // Get conversation by ID
		conv.POST("/", handler.CreateConversation)                     // Create a new conversation
		conv.PUT("/:id", handler.UpdateConversation)                   // Update a conversation
		conv.DELETE("/:id", handler.DeleteConversation)                // Delete a conversation
		conv.GET("/:id/message", handler.GetNextConversationMessage)   // Get next message in conversation
		conv.POST("/:id/message", handler.PostNextConversationMessage) // Post next message in conversation
	}

}
