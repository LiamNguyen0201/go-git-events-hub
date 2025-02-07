package models

// GitLabEvent represents the entire event structure
type GitLabEvent struct {
	ID             int64    `json:"id" gorm:"primaryKey"`
	ProjectID      int64    `json:"project_id" gorm:"index"`
	ActionName     string   `json:"action_name"`
	TargetID       *int64   `json:"target_id"`
	TargetIID      *int64   `json:"target_iid"`
	TargetType     *string  `json:"target_type"`
	AuthorID       int64    `json:"author_id"`
	TargetTitle    *string  `json:"target_title"`
	CreatedAt      string   `json:"created_at"`
	Author         Author   `json:"author" gorm:"serializer:json"`
	Imported       bool     `json:"imported"`
	ImportedFrom   string   `json:"imported_from"`
	PushData       PushData `json:"push_data" gorm:"serializer:json"`
	AuthorUsername string   `json:"author_username"`
}

// Author represents the author details
type Author struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	State     string `json:"state"`
	Locked    bool   `json:"locked"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}

// PushData represents push-related details
type PushData struct {
	CommitCount int     `json:"commit_count"`
	Action      string  `json:"action"`
	RefType     string  `json:"ref_type"`
	CommitFrom  *string `json:"commit_from"`
	CommitTo    string  `json:"commit_to"`
	Ref         string  `json:"ref"`
	CommitTitle string  `json:"commit_title"`
	RefCount    *int    `json:"ref_count"`
}
