package get_user_info

import (
	"context"
	oauth2service "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
	"log"
	"net/http"
)

func Email(client *http.Client) (string, error) {
	oauth2Service, err := oauth2service.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatal("Error creando servicio OAuth2: %w", err)
	}

	userInfo, serviceError := oauth2Service.Userinfo.Get().Do()
	if serviceError != nil {
		log.Fatal("error obteniendo información del usuario: %w", serviceError)
	}

	return userInfo.Email, nil
}
