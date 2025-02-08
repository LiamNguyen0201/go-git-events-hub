package dtos

type PullGitLabProjectRequest struct {
	ID int64 `json:"id" validate:"required"`
}
