package main

import (
	config "backend/internal/config"
	"log"
)

func main() {
	config.InitDB()                       // Initialize database connection
	router := config.IntiatializeRouter() // Initialize the router

	log.Println("Starting server on :8080")
	router.Run(":8080")

}
