package main

import (
	"log"
	"recordari/infrastructure/google_actions/google_calendar_service"
	"recordari/infrastructure/google_actions/google_get_calendar_events"
)

func main() {

	service := google_calendar_service.Get()

	if service == nil {
		log.Fatalf("Hubo un error en la lectura del fichero")
	}

	events := google_get_calendar_events.Get(service)

	if len(events) != 0 {

	}

}
