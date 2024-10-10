package v1

import (
	gin "github.com/gin-gonic/gin"
	"github.com/shrikanthcodes/butler-ai/internal/controller/handler"
	middleware "github.com/shrikanthcodes/butler-ai/internal/controller/http"
)

func TestRoutes(router *gin.Engine) {
	test := router.Group("/api/test")
	test.Use(middleware.GeneralAuthMiddleware())
	{
		// TEST: remove in next version
		test.GET("/healthcheck", handler.HealthCheck) // Health check
	}
}
