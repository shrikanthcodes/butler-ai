package v1

//import (
//	server "github.com/server-gonic/server"
//	"github.com/shrikanthcodes/butler-ai/internal/controller/handler"
//	middleware "github.com/shrikanthcodes/butler-ai/internal/controller/api"
//)
//
//func AdminRoutes(router *server.Engine) {
//	admin := router.Group("/api/admin")
//	admin.Use(middleware.AdminAuthMiddleware())
//	{
//		admin.PUT("/user/{id}/ban", handler.BanUser)             // Ban a user
//		admin.GET("/conversations", handler.GetAllConversations) // Get all conversations
//		admin.GET("/users", handler.GetAllUsers)                 // Get all users
//	}
//}
