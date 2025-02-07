package main

import (
	"git_events_hub/clients"
	"git_events_hub/configs"
	"git_events_hub/databases"
	"git_events_hub/handlers"
	"git_events_hub/middlewares"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Background function that runs every X seconds
func startTicker() {
	ticker := time.NewTicker(time.Duration(configs.PollInterval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Checking for new GitLab events...")
		events := clients.FetchGitLabEvents()

		for _, event := range events {
			if !databases.EventExists(event.ID) {
				log.Printf("New event found: ID=%d Action=%s", event.ID, event.ActionName)

				// Save to database
				databases.SaveEvent(event)

				// Forward to webhook
				clients.SendEventToWebhook(event)
			} else {
				log.Printf("Skipping existing event: ID=%d", event.ID)
			}
		}
	}
}

func main() {
	log.Println("Starting Git events hub Service...")

	databases.InitDB()

	// Start the ticker in a separate goroutine
	go startTicker()

	// Initialize Gin router
	r := gin.Default()

	r.Use(middlewares.PaginationMiddleware())

	// Add custom recovery middleware
	r.Use(middlewares.RecoveryWithLogger())

	r.GET("/events", handlers.GetEvents)
	r.GET("/events/:id", handlers.GetEventDetail)

	// Start server
	r.Run(":" + configs.Port)
}
