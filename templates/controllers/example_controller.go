package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ExampleController handles example-related requests
type ExampleController struct{}

// NewExampleController creates a new example controller
func NewExampleController() *ExampleController {
	return &ExampleController{}
}

// GetExamples returns a list of examples
func (ctrl *ExampleController) GetExamples(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get all examples",
	})
}

// GetExample returns a single example by ID
func (ctrl *ExampleController) GetExample(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Get example by ID",
		"id":     id,
	})
} 