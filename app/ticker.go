package main

import (
	"context"
	"git_events_hub/clients"
	"git_events_hub/configs"
	"git_events_hub/databases"
	"git_events_hub/models"
	"git_events_hub/utils"
	"os"
	"sync"
	"time"

	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
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
		if configs.EnableDistributedLock {
			ctx := context.Background()
			hostname, _ := os.Hostname()
			lockKey := "go-git-events-hub-ticker-distributed-lock"
			ttl := 60 * time.Second // Lock expiration time

			err := processWithLock(ctx, lockKey, ttl)
			if err != nil {
				utils.LogInfof("(startTicker) Worker %s: %v\n", hostname, err)
			} else {
				utils.LogInfof("(startTicker) Worker %s: Successfully processed\n", hostname)
			}
		} else {
			process()
		}
	}
}

// Acquire a distributed lock
func processWithLock(ctx context.Context, lockKey string, ttl time.Duration) error {
	// Redis client setup
	var redisClient = redis.NewClient(&redis.Options{
		Addr: configs.RedisHost, // Change if necessary
	})

	locker := redislock.New(redisClient)

	// Attempt to acquire lock
	lock, err := locker.Obtain(ctx, lockKey, ttl, nil)
	if err == redislock.ErrNotObtained {
		utils.LogInfof("(processWithLock) Failed to acquire lock: %s", lockKey)
		return err
	} else if err != nil {
		return err
	}

	defer lock.Release(ctx) // Ensure lock is released after processing

	// Critical Section (Execute your task safely)
	utils.LogInfo("(processWithLock) Lock acquired, processing task...")
	process()
	utils.LogInfof("(processWithLock) Task completed! Release lock!: %s", lockKey)

	return nil
}

func process() {
	utils.LogInfo("(process) Checking for new GitLab events...")
	time.Sleep(15 * time.Second) // Simulate processing
	events := clients.FetchGitLabEvents()

	if len(events) == 0 {
		utils.LogInfo("(process) No new events found.")
		return
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
			utils.LogInfof("(process) New event found: ID=%d Action=%s", event.ID, event.ActionName)

			// Save to database
			databases.SaveEvent(event)

			wg.Add(1)
			eventChannel <- event
		} else {
			utils.LogInfof("(process) Skipping existing event: ID=%d", event.ID)
		}
	}
	utils.LogInfo("(process) Done!")

	// Close channel after sending all events
	close(eventChannel)

	// Wait for all goroutines to finish
	wg.Wait()
}
