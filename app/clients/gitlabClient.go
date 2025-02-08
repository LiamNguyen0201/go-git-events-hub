package clients

import (
	"encoding/json"
	"git_events_hub/configs"
	"git_events_hub/models"
	"git_events_hub/utils"
	"io/ioutil"
	"net/http"
	"strconv"
)

func FetchGitLabEvents() []models.GitLabEvent {
	req, _ := http.NewRequest("GET", configs.GitLabAPIURL+"/api/v4/events", nil)
	req.Header.Set("PRIVATE-TOKEN", configs.GitLabToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode >= 400 {
		utils.LogInfo("(FetchGitLabEvents) Error fetching GitLab events", err)
		return nil
	}
	defer resp.Body.Close()
	utils.LogInfo("(FetchGitLabEvents) Status code: " + resp.Status)

	body, _ := ioutil.ReadAll(resp.Body)
	var events []models.GitLabEvent
	json.Unmarshal(body, &events)

	return events
}

func FetchGitLabProject(projectID int64) *models.GitLabProject {
	req, _ := http.NewRequest("GET", configs.GitLabAPIURL+"/api/v4/projects/"+strconv.FormatInt(projectID, 10), nil)
	req.Header.Set("PRIVATE-TOKEN", configs.GitLabToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode >= 400 {
		utils.LogInfo("(FetchGitLabProject) Error fetching GitLab project:", err)
		return nil
	}
	defer resp.Body.Close()
	utils.LogInfof("(FetchGitLabProject) Status code: %s", resp.Status)

	body, _ := ioutil.ReadAll(resp.Body)
	var project *models.GitLabProject
	json.Unmarshal(body, &project)

	return project
}
