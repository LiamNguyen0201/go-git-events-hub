package clients

import (
	"bytes"
	"encoding/json"
	"git_events_hub/configs"
	"git_events_hub/models"
	"log"
	"net/http"
)

// SendEventToWebhook forwards a GitLab event to an external endpoint
func SendEventToWebhook(event models.GitLabEvent) {
	eventJSON, _ := json.Marshal(event)

	req, _ := http.NewRequest("POST", configs.WebhookURL, bytes.NewBuffer(eventJSON))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Failed to send event:", err)
		return
	}
	defer resp.Body.Close()

	log.Println("Event forwarded successfully:", event.ID)
}
