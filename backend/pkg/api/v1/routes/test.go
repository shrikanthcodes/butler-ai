package routes

import (
	middleware "backend/pkg/api/middlewares"
	handler "backend/pkg/api/v1/handlers"

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
