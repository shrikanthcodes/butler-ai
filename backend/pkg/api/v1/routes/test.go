package routes

import (
	middleware "github.com/shrikanthcodes/butler-ai/backend/pkg/api/middlewares"
	handler "github.com/shrikanthcodes/butler-ai/backend/pkg/api/v1/handlers"

	gin "github.com/gin-gonic/gin"
)

func TestRoutes(router *gin.Engine) {
	test := router.Group("/api/test")
	test.Use(middleware.GeneralAuthMiddleware())
	{
		// TEST: remove in next version
		test.GET("/healthcheck", handler.HealthCheck) // Health check
	}
}
