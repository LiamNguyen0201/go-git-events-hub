package handlers

import (
	"git_events_hub/databases"
	"git_events_hub/dtos"
	"git_events_hub/models"
	"git_events_hub/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func CreateWorkflow(c *gin.Context) {
	var request dtos.WorkflowRequestDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.LogInfo("(CreateWorkflow) Bind JSON error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	workflow := models.Workflow{}
	copier.Copy(&workflow, &request)
	databases.CreateWorkflow(workflow)
	c.JSON(http.StatusCreated, "")
}

func GetWorkflow(c *gin.Context) {
	workflowID := c.Param("id")
	if workflowID == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Workflow not found"})
		return
	}

	workflow, err := databases.GetWorkflowByID(utils.StringToNumber(workflowID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Workflow not found"})
		return
	}

	c.JSON(http.StatusOK, workflow)
}

func GetWorkflows(c *gin.Context) {
	// Query parameters
	startDate := c.Query("start_date") // e.g., "2025-01-01"
	endDate := c.Query("end_date")     // e.g., "2025-02-01"
	page := c.GetInt("page")
	limit := c.GetInt("limit")

	events, total := databases.GetWorkflows(startDate, endDate, page, limit)

	c.JSON(http.StatusOK, gin.H{
		"total":   total,
		"page":    page,
		"limit":   limit,
		"results": events,
	})
}

func UpdateWorkflow(c *gin.Context) {
	workflowID := utils.StringToNumber(c.Param("id"))
	if !databases.DoesWorkflowExist(workflowID) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Workflow not found"})
		return
	}

	var request dtos.WorkflowRequestDTO

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.LogInfo("(CreateWorkflow) Bind JSON error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	workflow := models.Workflow{}
	copier.Copy(&workflow, &request)
	workflow.ID = workflowID
	databases.SaveWorkflow(workflow)
	c.JSON(http.StatusOK, workflow)
}
