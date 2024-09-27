package main

import (
	config "backend/internal/config"
	api "backend/pkg/api"
)

func main() {
	config.InitDB()  // Initialize database connection
	api.InitRoutes() // Initialize the routes

}
