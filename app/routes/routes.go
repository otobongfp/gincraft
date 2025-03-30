package routes

import (
	"app/controllers"
	"app/services"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all routes for the application
func RegisterRoutes(r *gin.Engine) {
	// Initialize services
	exampleService := services.NewExampleService()

	// Initialize controllers
	exampleController := controllers.NewExampleController(exampleService)

	// Define routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to app!",
		})
	})

	r.GET("/health", exampleController.HealthCheck)
	r.GET("/example", exampleController.GetExample)

	// API v1 group
	v1 := r.Group("/api/v1")
	{
		// Add your routes here
	}
} 