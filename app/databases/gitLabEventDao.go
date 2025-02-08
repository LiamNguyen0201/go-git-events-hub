package databases

import (
	"git_events_hub/models"
	"git_events_hub/utils"
	"strconv"
)

func EventExists(eventID int64) bool {
	var count int64
	db.Model(&models.GitLabEvent{}).Where("id = ?", eventID).Count(&count)
	return count > 0
}

func GetEventByID(eventID int64) (*models.GitLabEvent, error) {
	var event models.GitLabEvent

	err := db.First(&event, eventID).Error
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func GetEvents(projectID string, startDate string, endDate string, page int, limit int) (*[]models.GitLabEvent, int64) {
	utils.LogDebug("(GetEvents) Page: " + strconv.Itoa(page))
	utils.LogDebug("(GetEvents) Limit: " + strconv.Itoa(limit))

	var events []models.GitLabEvent
	var total int64

	offset := (page - 1) * limit // Pagination logic

	// Query builder
	query := db.Model(&models.GitLabEvent{})
	if projectID != "" {
		query = query.Where("project_id = ?", utils.StringToNumber(projectID))
	}
	if startDate != "" && endDate != "" {
		query = query.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}

	// Count total results
	query.Count(&total)

	// Fetch paginated results
	query.Limit(limit).Offset(offset).Find(&events)

	return &events, total // Return event
}

func SaveEvent(event models.GitLabEvent) {
	db.Create(&event)
}
