package domains

type Project struct {
	//TODO
}

type ProjectService interface {
	GetUserProjects(userID string) ([]Project, error)
	updateProject(project *Project) error
	deleteProject(project *Project) error
}

type ProjectRepository
