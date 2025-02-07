package clients

import (
	"encoding/json"
	"git_events_hub/configs"
	"git_events_hub/models"
	"git_events_hub/utils"
	"io/ioutil"
	"net/http"
)

// Fetch new events from GitLab
func FetchGitLabEvents() []models.GitLabEvent {
	req, _ := http.NewRequest("GET", configs.GitLabAPIURL, nil)
	req.Header.Set("PRIVATE-TOKEN", configs.GitLabToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		utils.LogInfo("Error fetching GitLab events:", err)
		return nil
	}
	defer resp.Body.Close()
	utils.LogInfo("Status code: " + resp.Status)

	body, _ := ioutil.ReadAll(resp.Body)
	var events []models.GitLabEvent
	json.Unmarshal(body, &events)

	return events
}
