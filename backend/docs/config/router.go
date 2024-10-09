package config

import (
	cors "github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
	middleware "github.com/shrikanthcodes/butler-ai/backend/docs"
)

func IntiatializeRouter() *gin.Engine {
	router := gin.Default()

	// Add CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // React frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	router.Use(middleware.LoggingMiddleware())

	return router
}
