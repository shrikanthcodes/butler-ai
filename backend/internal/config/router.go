package config

import (
	middleware "backend/pkg/api/middlewares"

	cors "github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
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
