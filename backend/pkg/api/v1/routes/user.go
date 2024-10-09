package routes

// import (
// 	middleware "backend/pkg/api/middlewares"
// 	handler "backend/pkg/api/v1/handlers"

// 	gin "github.com/gin-gonic/gin"
// )

// func UserRoutes(router *gin.Engine) {
// 	user := router.Group("/api/user")
// 	user.Use(middleware.GeneralAuthMiddleware())
// 	{
// 		user.GET("/{id}", handler.GetUserByID)                        // Get user by ID
// 		user.POST("/", handler.CreateUser)                            // Create a new user
// 		user.PUT("/{id}", handler.UpdateUser)                         // Update a user
// 		user.DELETE("/{id}", handler.DeleteUser)                      // Delete a user
// 		user.GET("/{id}/conversations", handler.GetUserConversations) // Get conversations associated by user ID
// 		user.GET("/{id}/profile", handler.GetUserProfile)             // Get user's Profile by user ID
// 		user.PUT("/{id}/profile", handler.UpdateUserProfile)          // Update user's Profile by user ID
// 		user.GET("/{id}/health", handler.GetUserHealth)               // Get user's Health by user ID
// 		user.PUT("/{id}/health", handler.UpdateUserHealth)            // Update user's Health by user ID
// 		user.GET("/{id}/diet", handler.GetUserDiet)                   // Get user's Diet by user ID
// 		user.PUT("/{id}/diet", handler.UpdateUserDiet)                // Update user's Diet by user ID
// 		user.GET("/{id}/inventory", handler.GetUserInventory)         // Get user's Inventory by user ID
// 		user.PUT("/{id}/inventory", handler.UpdateUserInventory)      // Update user's Inventory by user ID
// 		user.GET("/{id}/goal", handler.GetUserGoal)                   // Get user's Goal by user ID
// 		user.PUT("/{id}/goal", handler.UpdateUserGoal)                // Update user's Goal by user ID
// 		user.GET("/{id}/llm", handler.GetUserLLM)                     // Get user's LLM by user ID
// 		user.PUT("/{id}/llm", handler.UpdateUserLLM)                  // Update user's LLM by user ID
// 		user.GET("/{id}/script", handler.GetUserScript)               // Get user's Script by user ID
// 		user.PUT("/{id}/script", handler.UpdateUserScript)            // Update user's Script by user ID
// 		user.GET("/{id}/shopping", handler.GetUserShopping)           // Get user's Shopping by user ID
// 		user.PUT("/{id}/shopping", handler.UpdateUserShopping)        // Update user's Shopping by user ID
// 		user.GET("/{id}/mealchoices", handler.GetUserMealChoices)     // Get user's MealChoices by user ID
// 		user.PUT("/{id}/mealchoices", handler.UpdateUserMealChoices)  // Update user's MealChoices by user ID
// 		user.GET("/{id}/choices", handler.GetAllChoices)              // Get all choices
// 		user.PUT("/{id}/choices", handler.UpdateChoices)              // Update choices
// 		user.GET("/{id}/writeup", handler.GetUserWriteUp)             // Get user's WriteUp by user ID
// 		user.POST("/{id}/writeup", handler.UpdateUserWriteUp)         // Update user's WriteUp by user ID
// 	}
// }
