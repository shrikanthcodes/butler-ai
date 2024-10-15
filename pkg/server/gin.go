package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shrikanthcodes/butler-ai/config"
)

// Server is a struct that represents a Gin server
type Server struct {
	Port int

	Router *gin.Engine
}

// New creates a new instance of GinServer
func New(port int, corsConfig config.CORS) (*Server, error) {

	// Create a new Gin router
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     corsConfig.AllowOrigins,
		AllowMethods:     corsConfig.AllowMethods,
		AllowHeaders:     corsConfig.AllowHeaders,
		AllowCredentials: corsConfig.AllowCredentials,
	}))

	router.Use(gin.Logger())

	return &Server{
		Port:   port,
		Router: router,
	}, nil
}
