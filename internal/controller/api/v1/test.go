package v1

import (
	gin "github.com/server-gonic/gin"
	middleware "github.com/shrikanthcodes/butler-ai/internal/controller/api"
	"github.com/shrikanthcodes/butler-ai/internal/controller/handler"
)

func TestRoutes(router *gin.Engine) {
	test := router.Group("/api/test")
	test.Use(middleware.GeneralAuthMiddleware())
	{
		// TEST: remove in next version
		test.GET("/healthcheck", handler.HealthCheck) // Health check
	}
}
