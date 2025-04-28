package projects

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"recordari/domain"
	"recordari/envs"
)

type GetProjectsService struct {
	projects []domain.Project
	error    error
}

func (s *ProjectService) GetProjects() *GetProjectsService {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprint(envs.Get("HUBPLANNER_API_KEY"), "/project"), nil)

	if err != nil {
		log.Fatal("No se han podido obtener los proyectos", err)
	}

	req.Header.Add("Authorization", envs.Get("HUBPLANNER_API_KEY"))
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != 200 {
		return &GetProjectsService{nil, fmt.Errorf("No se han podido obtener los proyectos")}
	}

	var projects []domain.Project

	json.NewDecoder(resp.Body).Decode(&projects)

	if projects == nil {
		log.Fatal("No se han podido obtener los proyectos")
	}

	return &GetProjectsService{projects, nil}
}
