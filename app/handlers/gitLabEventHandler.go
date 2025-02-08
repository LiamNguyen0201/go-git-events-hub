package handlers

import (
	"git_events_hub/databases"
	"git_events_hub/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
	// Query parameters
	projectID := c.Query("project_id")
	startDate := c.Query("start_date") // e.g., "2025-01-01"
	endDate := c.Query("end_date")     // e.g., "2025-02-01"
	page := c.GetInt("page")
	limit := c.GetInt("limit")

	events, total := databases.GetEvents(projectID, startDate, endDate, page, limit)

	c.JSON(http.StatusOK, gin.H{
		"total":   total,
		"page":    page,
		"limit":   limit,
		"results": events,
	})
}

func GetEventDetail(c *gin.Context) {
	eventID := c.Param("id")
	if eventID == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	event, err := databases.GetEventByID(utils.StringToNumber(eventID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, event)
}
