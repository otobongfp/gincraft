package controllers

import (
	"net/http"

	"app/services"

	"github.com/gin-gonic/gin"
)

// ExampleController handles example-related requests
type ExampleController struct {
	exampleService *services.ExampleService
}

// NewExampleController creates a new example controller
func NewExampleController(exampleService *services.ExampleService) *ExampleController {
	return &ExampleController{
		exampleService: exampleService,
	}
}

// GetExamples returns a list of examples
func (ctrl *ExampleController) GetExamples(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get all examples",
	})
}

// GetExample returns a single example by ID
func (c *ExampleController) GetExample(ctx *gin.Context) {
	result, err := c.exampleService.GetExample()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (c *ExampleController) HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
} 