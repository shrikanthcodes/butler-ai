package v1

import (
	gin "github.com/gin-gonic/gin"
	"github.com/shrikanthcodes/butler-ai/backend/internal/controller/handler"
)

func AuthRoutes(router *gin.Engine) {
	auth := router.Group("/api")
	//auth.Use(middleware.GeneralAuthMiddleware())
	{
		auth.POST("/login", handler.Login)                       // Login
		auth.POST("/register", handler.Register)                 // Register
		auth.POST("/logout", handler.Logout)                     // Logout
		auth.POST("/token/refresh", handler.RefreshToken)        // Refresh token JWT
		auth.POST("/user/reset_password", handler.ResetPassword) // Reset password
	}
}
