package calendars

import (
	"context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"recordari/services/google/credentials"
	"recordari/services/google/generate_token"
)

func GetGoogleCalendarService() (*calendar.Service, *http.Client) {
	googleCredentials := credentials.GetGoogleCredentials("Ha ocurrido un error al leer el fichero de credenciales.")

	// Readonly calendar permission
	config, _ := google.ConfigFromJSON(googleCredentials,
		"https://www.googleapis.com/auth/userinfo.email",
		calendar.CalendarReadonlyScope)

	token, _ := generate_token.Google(config)
	client := config.Client(context.Background(), token)
	service, err := calendar.NewService(context.Background(), option.WithHTTPClient(client))

	if err != nil {
		log.Fatalf("❌ No se pudo crear el servicio de Calendar: %v", err)
	}

	return service, client
}
