package clients

import (
	"bytes"
	"encoding/json"
	"git_events_hub/configs"
	"git_events_hub/models"
	"log"
	"net/http"
	"time"
)

// Max retries for failed requests
const maxRetries = 3

// SendEventToWebhook forwards a GitLab event to an external endpoint
func SendEventToWebhook(event models.GitLabEvent, retryCount int) {
	if retryCount >= maxRetries {
		log.Printf("Failed to send event ID=%d after %d retries\n", event.ID, maxRetries)
		return
	}

	eventJSON, _ := json.Marshal(event)
	req, _ := http.NewRequest("POST", configs.WebhookURL, bytes.NewBuffer(eventJSON))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)

	if err != nil || resp.StatusCode >= 400 {
		log.Printf("Error sending event ID=%d (Attempt %d): %v\n", event.ID, retryCount+1, err)

		// Exponential backoff before retrying
		time.Sleep(time.Duration(2<<retryCount) * time.Second)

		// Retry recursively
		SendEventToWebhook(event, retryCount+1)
		return
	}
	defer resp.Body.Close()

	log.Printf("Event ID=%d sent successfully\n", event.ID)
}
