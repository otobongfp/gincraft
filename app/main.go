package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"PROJECT_NAME/controllers"
	"PROJECT_NAME/routes"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Initialize routes
	routes.SetupRoutes(router)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
} 