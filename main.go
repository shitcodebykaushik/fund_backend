package main

import (
	//"fund/controllers"
	"fund/db"
	"fund/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize MongoDB connection
	db.Init()

	// Initialize Gin router
	router := gin.Default()

	// Set up routes
	routes.AuthRoutes(router)

	// Start the server
	router.Run(":8080")
}
