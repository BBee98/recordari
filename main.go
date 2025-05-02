package main

import (
	"recordari/services/google/calendars"
	"recordari/services/google/events"
	"recordari/services/google/get_user_info"
	"recordari/services/hubplanner"
	"recordari/ui"
)

func main() {

	service, client := calendars.GetGoogleCalendarService()
	calendarEvents := events.GetGoogleCalendarEvents(service)
	userEmail, emailErr := get_user_info.Email(client)
	resources, resourcesErr := hubplanner.GetResources()
	projects, projectsErr := hubplanner.GetProjects()

	if emailErr != nil {
		println("No se ha podido obtener el email", emailErr)
	}

	if resourcesErr != nil {
		println("No se han podido obtener los recursos")
	}

	if projectsErr != nil {
		println("No se han podido obtener los proyectos")
	}

	resourceId := hubplanner.GetResourceIdFromEmail(resources, userEmail)

	// projectId := hubplanner.GetResourcesFromProjectsWithId(projects, resourceId)

	hubplanner.RecordUser(resourceId, userEmail)
	userProjects := hubplanner.GetProjectsFromResourceId(projects, resourceId)

	println(calendarEvents)

	fyneUI := ui.FyneUI{}
	fyneUI.Mount().ListProjects(userProjects).Show()
}
