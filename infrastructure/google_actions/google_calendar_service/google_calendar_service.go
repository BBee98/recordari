package google_calendar_service

import (
	"context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"log"
	"recordari/infrastructure/google_actions/google_credentials"
	"recordari/infrastructure/google_actions/google_generate_token"
)

func Get() *calendar.Service {
	googleCredentials, credentialsError := google_credentials.Get()

	if credentialsError != nil {
		println(credentialsError)
		log.Fatalf("No se han podido obtener las credenciales de Google")
	}
	// Readonly calendar permission
	config, _ := google.ConfigFromJSON(googleCredentials, calendar.CalendarReadonlyScope)

	token, _ := google_generate_token.Generate(config)
	client := config.Client(context.Background(), token)
	service, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))

	if err != nil {
		log.Fatalf("❌ No se pudo crear el servicio de Calendar: %v", err)
	}

	return service
}
