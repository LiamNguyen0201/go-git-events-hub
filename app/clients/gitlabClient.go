package clients

import (
	"git_events_hub/configs"
	"git_events_hub/models"
	"git_events_hub/utils"
	"strconv"

	"resty.dev/v3"
)

func FetchGitLabEvents() []models.GitLabEvent {
	var events []models.GitLabEvent

	c := resty.New()
	defer c.Close()

	resp, err := c.R().
		SetHeader("PRIVATE-TOKEN", configs.GitLabToken).
		SetResult(events).
		Get(configs.GitLabAPIURL + "/api/v4/events")
	utils.LogInfof("(FetchGitLabEvents) Status code: %d", resp.StatusCode())

	if err != nil || resp.StatusCode() >= 400 {
		utils.LogInfo("(FetchGitLabEvents) Error fetching GitLab events", err)
		return nil
	}

	return events
}

func FetchGitLabProject(projectID int64) *models.GitLabProject {
	utils.LogDebugf("(FetchGitLabProject) GitLab project : %d", projectID)

	var project *models.GitLabProject

	c := resty.New()
	defer c.Close()

	resp, err := c.R().
		SetHeader("PRIVATE-TOKEN", configs.GitLabToken).
		SetResult(project).
		Get(configs.GitLabAPIURL + "/api/v4/projects/" + strconv.FormatInt(projectID, 10))
	utils.LogInfof("(FetchGitLabProject) Status code: %d", resp.StatusCode())

	if err != nil || resp.StatusCode() >= 400 {
		utils.LogInfof("(FetchGitLabProject) Error fetching GitLab project : %d, detail: %s", projectID, err)
		return nil
	}

	return project
}
