package models

import "time"

type ProcessorType string

const (
	NOTIFY_JENKIN_PROCESSOR      ProcessorType = "NOTIFY_JENKIN_PROCESSOR"
	NOTIFY_SLACK_PROCESSOR       ProcessorType = "NOTIFY_SLACK_PROCESSOR"
	PULL_GIT_LAB_EVENT_PROCESSOR ProcessorType = "PULL_GIT_LAB_EVENT_PROCESSOR"
	SAVE_GIT_LAB_EVENT_PROCESSOR ProcessorType = "SAVE_GIT_LAB_EVENT_PROCESSOR"
)

type Workflow struct {
	ID           int64          `json:"id" gorm:"primaryKey"`
	Name         string         `json:"name"`
	Cron         string         `json:"cron"`
	HttpEndpoint string         `json:"http_endpoint"`
	Nodes        []WorkflowNode `json:"nodes" gorm:"serializer:json"`
	IsActive     bool           `json:"is_active"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

type WorkflowRunResult struct {
	ID         int64           `json:"id" gorm:"primaryKey"`
	WorkflowID int64           `json:"workflow_id" gorm:"index"`
	Results    []NodeRunResult `json:"results" gorm:"serializer:json"`
	IsSuccess  bool            `json:"is_success"`
	StartedAt  time.Time       `json:"started_at"`
	EndedAt    time.Time       `json:"ended_at"`
}

// supported struct: WorkflowNode represents the process of a node in a workflow
type WorkflowNode struct {
	ID             int64         `json:"id"`
	PreviousNodeID int64         `json:"previous_node_id"`
	NextNodeID     int64         `json:"next_node_id"`
	Description    string        `json:"name"`
	ProcessorType  ProcessorType `json:"processor_type"`
	ProcessorData  string        `json:"processor_data"`
}

// supported struct: NodeRunResult represents the run of a workflow node details
type NodeRunResult struct {
	WorkflowNodeID int64  `json:"workflow_node_id"`
	Input          string `json:"input"`
	Output         string `json:"output"`
	IsSuccess      bool   `json:"is_success"`
}

// supported struct
type NotifySlackProcessor struct {
	WebHookUrl string `json:"web_hook_url"`
}

// supported struct
type PullGitLabEventProcessor struct {
	PersonalAccessTokenID int64   `json:"personal_access_token_id"`
	ProjectIDs            []int64 `json:"project_ids"`
}
