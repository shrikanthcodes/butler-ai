package api

import (
	"log"

	config "backend/internal/config"
	middleware "backend/pkg/api/middlewares"
)

func InitRoutes() {
	// Set the routes
	router := config.IntiatializeRouter() // Initialize the router
	// TEST: remove in next version
	router.GET("/api/health", middleware.HealthCheck)
	router.POST("/api/user/random", middleware.CreateRandomUser)

	log.Println("Starting server on :8080")
	router.Run(":8080")

	// router.POST("/api/login", middleware.Login)
	// router.POST("/api/register", middleware.Register)
	// router.POST("/api/logout", middleware.Logout)

	// Conversation routes
	// router.GET("/api/conversation/{id}", middleware.GetConversationByID) // Get conversation by ID
	// router.GET("/api/conversation", middleware.GetAllConversations)         // Get all conversations
	// router.POST("/api/conversation", middleware.CreateConversation)      // Create a new conversation
	// router.PUT("/api/conversation/{id}", middleware.UpdateConversation)   // Update a conversation
	// router.DELETE("/api/conversation/{id}", middleware.DeleteConversation) // Delete a conversation
	// router.GET("/api/conversation/{id}/message", middleware.GetNextConversationMessage) // Get next message in conversation
	// router.POST("/api/conversation/{id}/message", middleware.PostNextConversationMessage) // Post next message in conversation

	// Other routes
	// router.GET("/api/user/{id}", middleware.GetUserByID) // Get user by ID
	// router.GET("/api/user", middleware.GetAllUsers)         // Get all users
	// router.POST("/api/user", middleware.CreateUser)      // Create a new user
	// router.PUT("/api/user/{id}", middleware.UpdateUser)   // Update a user
	// router.DELETE("/api/user/{id}", middleware.DeleteUser) // Delete a user
	// router.GET("/api/user/{id}/conversations", middleware.GetUserConversations) // Get conversations associated by user ID
	// router.GET("/api/user/{id}/profile", middleware.GetUserProfile) // Get user's Profile by user ID
	// router.PUT("/api/user/{id}/profile", middleware.UpdateUserProfile) // Update user's Profile by user ID
	// router.GET("/api/user/{id}/health", middleware.GetUserHealth) // Get user's Health by user ID
	// router.PUT("/api/user/{id}/health", middleware.UpdateUserHealth) // Update user's Health by user ID	// router.GET("/api/user/{id}/diet", middleware.GetUserDiet) // Get user's Diet by user ID
	// router.PUT("/api/user/{id}/diet", middleware.UpdateUserDiet) // Update user's Diet by user ID
	// router.GET("/api/user/{id}/inventory", middleware.GetUserInventory) // Get user's Inventory by user ID
	// router.PUT("/api/user/{id}/inventory", middleware.UpdateUserInventory) // Update user's Inventory by user ID
	// router.GET("/api/user/{id}/goal", middleware.GetUserGoal) // Get user's Goal by user ID
	// router.PUT("/api/user/{id}/goal", middleware.UpdateUserGoal) // Update user's Goal by user ID
	// router.GET("/api/user/{id}/llm", middleware.GetUserLLM) // Get user's LLM by user ID
	// router.PUT("/api/user/{id}/llm", middleware.UpdateUserLLM) // Update user's LLM by user ID
	// router.GET("/api/user/{id}/script", middleware.GetUserScript) // Get user's Script by user ID
	// router.PUT("/api/user/{id}/script", middleware.UpdateUserScript) // Update user's Script by user ID
	// router.GET("/api/user/{id}/shopping", middleware.GetUserShopping) // Get user's Shopping by user ID
	// router.PUT("/api/user/{id}/shopping", middleware.UpdateUserShopping) // Update user's Shopping by user ID
	// router.GET("/api/user/{id}/mealchoices", middleware.GetUserMealChoices) // Get user's MealChoices by user ID
	// router.PUT("/api/user/{id}/mealchoices", middleware.UpdateUserMealChoices) // Update user's MealChoices by user ID
	// router.GET("/api/choices", middleware.GetAllChoices) // Get all choices
	// router.PUT("/api/choices", middleware.UpdateChoices) // Update choices

	// router.GET("/api/user/{id}/writeup", middleware.GetUserWriteUp) // Get user's WriteUp by user ID
	// router.POST("/api/user/{id}/writeup", middleware.UpdateUserWriteUp) // Update user's WriteUp by user ID
}
