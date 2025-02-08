package databases

import "git_events_hub/models"

func GetProjectByID(projectID int64) (*models.GitLabProject, error) {
	var project models.GitLabProject

	err := db.First(&project, projectID).Error
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func SaveProject(project models.GitLabProject) {
	db.Create(&project)
}
