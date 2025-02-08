package handlers

import (
	"git_events_hub/clients"
	"git_events_hub/databases"
	"git_events_hub/dtos"
	"git_events_hub/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PullProject(c *gin.Context) {
	var request dtos.PullGitLabProjectRequest

	// Bind JSON request
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.LogInfo("(PullProject) Bind JSON error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	project := clients.FetchGitLabProject(request.ID)
	databases.SaveProject(*project)

	c.JSON(http.StatusOK, project)
}

func GetProjectDetail(c *gin.Context) {
	projectID := c.Param("id")
	if projectID == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	event, err := databases.GetProjectByID(utils.StringToNumber(projectID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}

	c.JSON(http.StatusOK, event)
}
