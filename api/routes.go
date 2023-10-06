package api

import (
	"github.com/gin-gonic/gin"
)

// InitRoutes initializes all API routes.
func InitRoutes(router *gin.Engine) {
	// Create a group for API routes (e.g., /api/v1)
	apiGroup := router.Group("/api/v1")

	// Define routes for different API endpoints
	apiGroup.POST("/libraries", createLibrary)
	apiGroup.GET("/libraries", getLibraries)
	apiGroup.POST("/books", addBook)
	apiGroup.GET("/books", getBooks)
	apiGroup.POST("/requests", raiseRequest)
	apiGroup.GET("/requests", listRequests)
	apiGroup.PUT("/requests/approve/:reqID", approveRequest)
	apiGroup.PUT("/requests/reject/:reqID", rejectRequest)
}
