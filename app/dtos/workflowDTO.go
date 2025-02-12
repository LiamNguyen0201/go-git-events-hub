package dtos

import (
	"git_events_hub/models"
)

type WorkflowRequestDTO struct {
	Name         string                `json:"name" validate:"required"`
	Cron         string                `json:"cron"`
	HttpEndpoint string                `json:"http_endpoint"`
	Nodes        []models.WorkflowNode `json:"nodes"`
	IsActive     bool                  `json:"is_active" validate:"required"`
}
