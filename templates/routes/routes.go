package routes

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all routes for the application
func RegisterRoutes(r *gin.Engine) {
	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// API v1 group
	v1 := r.Group("/api/v1")
	{
		// Add your routes here
	}
} 