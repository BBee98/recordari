package projects

import (
	"net/http"
	"recordari/envs"
)

type ProjectService struct {
	client  *http.Client
	baseURL string
}

func NewProjectService() *ProjectService {

	return &ProjectService{
		client:  &http.Client{},
		baseURL: envs.Get("HUBPLANNER_BASE_URL"),
	}
}
