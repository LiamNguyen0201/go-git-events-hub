package models

// GitLabEvent represents the entire event structure
type GitLabProject struct {
	ID            int64  `json:"id" gorm:"primaryKey"`
	Name          string `json:"name"`
	SshUrlToRepo  string `json:"ssh_url_to_repo"`
	HttpUrlToRepo string `json:"http_url_to_repo"`
}
