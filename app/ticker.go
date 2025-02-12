package main

import (
	"git_events_hub/clients"
	"git_events_hub/configs"
	"git_events_hub/databases"
	"git_events_hub/models"
	"git_events_hub/utils"
	"sync"
	"time"
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
