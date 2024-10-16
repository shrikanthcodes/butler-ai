package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/shrikanthcodes/butler-ai/config"
	"github.com/shrikanthcodes/butler-ai/pkg/logger"
	"github.com/shrikanthcodes/butler-ai/pkg/server"
)

func RegisterRoutes(port int, corsConfig config.CORS, log *logger.Logger) (*gin.Engine, error) {

	// Initialize API server
	_, err := server.New(port, corsConfig)
	if err != nil {
		log.Fatal("Failed to initialize Gin server", err)
	}
	// Register routes
	//_, err = api.RegisterRoutes(ginServer.Router)
	//if err != nil {
	//	log.Fatal("Failed to register routes", err)
	//}
	return nil, nil
}

// Close closes the gin server
func Close() error {
	// Close the gin server
	//err := ginServer.Close()
	//if err != nil {
	//	return err
	//}
	return nil
}
