package databases

import (
	"git_events_hub/models"
	"git_events_hub/utils"
	"strconv"
)

func GetWorkflowByID(workflowID int64) (*models.Workflow, error) {
	var workflow models.Workflow

	err := db.First(&workflow, workflowID).Error
	if err != nil {
		return nil, err
	}

	return &workflow, nil
}

func GetWorkflows(startDate string, endDate string, page int, limit int) (*[]models.Workflow, int64) {
	utils.LogDebug("(GetWorkflows) Page: " + strconv.Itoa(page))
	utils.LogDebug("(GetWorkflows) Limit: " + strconv.Itoa(limit))

	var workflows []models.Workflow
	var total int64

	offset := (page - 1) * limit // Pagination logic

	// Query builder
	query := db.Model(&models.Workflow{})
	if startDate != "" && endDate != "" {
		query = query.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}

	// Count total results
	query.Count(&total)

	// Fetch paginated results
	query.Limit(limit).Offset(offset).Find(&workflows)

	return &workflows, total // Return event
}

func DoesWorkflowExist(workflowID int64) bool {
	var count int64
	db.Model(&models.Workflow{}).Where("id = ?", workflowID).Count(&count)
	return count > 0
}

func CreateWorkflow(workflow models.Workflow) {
	db.Create(&workflow)
}

func SaveWorkflow(workflow models.Workflow) {
	db.Save(&workflow)
}
