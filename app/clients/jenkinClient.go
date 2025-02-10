package clients

import (
	"git_events_hub/configs"
	"git_events_hub/models"
	"git_events_hub/utils"
	"time"

	"resty.dev/v3"
)

// Max retries for failed requests
const maxRetries = 3

// SendEventToWebhook forwards a GitLab event to an external endpoint
func SendEventToWebhook(event models.GitLabEvent, retryCount int) {
	if retryCount >= maxRetries {
		utils.LogInfof("Failed to send event ID=%d after %d retries\n", event.ID, maxRetries)
		return
	}

	c := resty.New()
	defer c.Close()

	resp, err := c.R().
		SetBody(event).
		Post(configs.WebhookURL)

	if err != nil || resp.StatusCode() >= 400 {
		utils.LogInfof("Error sending event ID=%d (Attempt %d): %v\n", event.ID, retryCount+1, err)

		// Exponential backoff before retrying
		time.Sleep(time.Duration(2<<retryCount) * time.Second)

		// Retry recursively
		SendEventToWebhook(event, retryCount+1)
		return
	}
	defer resp.Body.Close()

	utils.LogInfof("Event ID=%d sent successfully\n", event.ID)
}
