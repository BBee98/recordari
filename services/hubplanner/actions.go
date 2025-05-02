package hubplanner

import (
	"encoding/json"
	"log"
	"recordari/domain"
	"recordari/services/file"
)

func GetProjects() ([]domain.Project, error) {
	resp, err := DoRequest("GET", "/project", nil)
	if err != nil {
		return nil, err
	}

	var projects []domain.Project

	errDecoding := json.NewDecoder(resp.Body).Decode(&projects)
	defer resp.Body.Close()

	if errDecoding != nil {
		log.Fatal("No se pudo decodificar los proyectos", errDecoding)
	}

	return projects, nil
}

func GetResourcesFromProjectsWithId(projects []domain.Project, resourceId string) string {
	for _, project := range projects {
		for _, resource := range project.Resources {
			if resource == resourceId {
				return project.ID
			}
		}
	}
	return ""
}

func GetResources() (domain.Resource, error) {
	resp, err := DoRequest("GET", "/resource", nil)
	if err != nil {
		log.Fatal("No se pudo obtener los recursos ", err)
	}
	defer resp.Body.Close()

	var resources domain.Resource

	errDecoding := json.NewDecoder(resp.Body).Decode(&resources)

	if errDecoding != nil {
		log.Fatal("No se pudo decodificar los recursos ", errDecoding)
	}

	return resources, nil
}

func GetResourceIdFromEmail(resources domain.Resource, email string) string {
	for _, resource := range resources {
		if string(resource.Email) == email {
			return resource.ID
		}
	}
	return ""
}

func RecordUser(resourceId string, email string) {
	user := &domain.User{
		Id:    resourceId,
		Email: email,
	}

	userBytes, err := json.Marshal(user)

	if err != nil {
		log.Fatal("Error al serializar el usuario", err)
	}

	file.CreateAndWriteFile("user.json", userBytes)
}

func RecordMeetings(projects []domain.Project, events []domain.Events, resourceId string) {

	thereAreEvents := len(events) != 0
	if thereAreEvents {

	}
}
