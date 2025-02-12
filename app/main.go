package main

import (
	"git_events_hub/clients"
	"git_events_hub/configs"
	"git_events_hub/databases"
	"git_events_hub/handlers"
	"git_events_hub/middlewares"
	"git_events_hub/models"
	"git_events_hub/utils"
	"log"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

const workerCount = 5 // Number of concurrent workers

// ProcessEvent wraps SendEventToWebhook for concurrency
func ProcessEvent(event models.GitLabEvent) {
	clients.SendEventToWebhook(event, 0) // Start with 0 retries
}

// Background function that runs every X seconds
func startTicker() {
	ticker := time.NewTicker(time.Duration(configs.PollInterval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		utils.LogInfo("Checking for new GitLab events...")
		events := clients.FetchGitLabEvents()

		if len(events) == 0 {
			utils.LogInfo("No new events found.")
			continue
		}

		// Use a worker pool to process events concurrently
		var wg sync.WaitGroup
		eventChannel := make(chan models.GitLabEvent, len(events))

		// Start worker goroutines
		for i := 0; i < workerCount; i++ {
			go func() {
				for event := range eventChannel {
					ProcessEvent(event)
					wg.Done()
				}
			}()
		}

		// Send events to the channel
		for _, event := range events {
			if !databases.EventExists(event.ID) {
				utils.LogInfof("New event found: ID=%d Action=%s", event.ID, event.ActionName)

				// Save to database
				databases.SaveEvent(event)

				wg.Add(1)
				eventChannel <- event
			} else {
				utils.LogInfof("Skipping existing event: ID=%d", event.ID)
			}
		}

		// Close channel after sending all events
		close(eventChannel)

		// Wait for all goroutines to finish
		wg.Wait()
	}
}

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
