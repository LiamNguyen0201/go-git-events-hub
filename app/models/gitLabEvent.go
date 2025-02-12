package models

import "time"

// GitLabEvent represents the entire event structure
type GitLabEvent struct {
	ID             int64    `json:"id" gorm:"primaryKey"`
	ProjectID      int64    `json:"project_id" gorm:"index"`
	ActionName     string   `json:"action_name"`
	Author         Author   `json:"author" gorm:"serializer:json"`
	AuthorUsername string   `json:"author_username"`
	PushData       PushData `json:"push_data" gorm:"serializer:json"`

	Imported     bool    `json:"imported"`
	ImportedFrom string  `json:"imported_from"`
	TargetID     *int64  `json:"target_id"`
	TargetIID    *int64  `json:"target_iid"`
	TargetTitle  *string `json:"target_title"`
	TargetType   *string `json:"target_type"`

	CreatedAt time.Time `json:"created_at"`
}

// supported struct: Author represents the author details
type Author struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	State     string `json:"state"`
	Locked    bool   `json:"locked"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}

// supported struct: PushData represents push-related details
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
