package dtos

type GitLabProjectRequestDTO struct {
	ID int64 `json:"id" validate:"required"`
}
