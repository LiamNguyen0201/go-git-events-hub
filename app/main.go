package main

import (
	"git_events_hub/configs"
	"git_events_hub/databases"
	"git_events_hub/handlers"
	"git_events_hub/middlewares"
	"git_events_hub/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting Git events hub Service...")

	// ** Setup Logrus **
	logger := utils.GetLogger()

	// ** Setup DB **
	databases.InitDB(logger)

	// Start the ticker in a separate goroutine
	if configs.EnableTicker {
		go startTicker()
	}

	// Initialize Gin router
	r := gin.Default()

	// Set request logger
	r.Use(middlewares.LoggerMiddleware(logger))

	// Set default page and limit in case of missing
	r.Use(middlewares.PaginationMiddleware())

	// Add custom recovery middleware
	r.Use(middlewares.RecoveryWithLogger(logger))

	r.GET("/api/gitlab/events", handlers.GetEvents)
	r.GET("/api/gitlab/events/:id", handlers.GetEventDetail)

	r.POST("/api/gitlab/projects", handlers.PullProject)
	r.GET("/api/gitlab/projects/:id", handlers.GetProjectDetail)

	// Workflow Routes
	r.GET("/api/workflows", handlers.GetWorkflows)
	r.POST("/api/workflows", handlers.CreateWorkflow)
	r.GET("/api/workflows/:id", handlers.GetWorkflow)
	r.PUT("/api/workflows/:id", handlers.UpdateWorkflow)

	// Start server
	logger.Info("Starting Gin server on :8080")
	if err := r.Run(":" + configs.Port); err != nil {
		logger.Fatal("Failed to start server:", err)
	}
}
