package google_generate_token

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/browser"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
	"time"

	"context"
)

const TokenFile = "token.json"

func Generate(config *oauth2.Config) (*oauth2.Token, error) {

	codeChan := make(chan string)

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			http.Error(w, "Código no recibido", http.StatusBadRequest)
			return
		}
		fmt.Fprint(w, "✅ Código recibido. Ya puedes cerrar esta pestaña.")
		codeChan <- code
	})
	go func() {
		log.Println("🌐 Servidor escuchando en http://localhost:8080/callback")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("❌ Error en el servidor: %v", err)
		}
	}()

	time.Sleep(1 * time.Second)

	tokenData, errReadingFile := os.ReadFile(TokenFile)

	// Token not found
	if errReadingFile != nil {
		fmt.Println("No se ha encontrado el fichero de token. Vamos a generar un nuevo")
		authURL := config.AuthCodeURL(
			"state-token",
			oauth2.AccessTypeOffline,
			oauth2.SetAuthURLParam("prompt", "consent"))
		browser.OpenURL(authURL)

		authCode := <-codeChan

		file, err := os.Create("token.json")

		if err != nil {
			log.Fatalf("Ha ocurrido un error al crear el fichero de token: %v", err)
		}

		token, _ := config.Exchange(context.Background(), authCode)
		json.NewEncoder(file).Encode(token)
		return token, nil
	}

	var token oauth2.Token
	jsonErr := json.Unmarshal(tokenData, &token)
	if jsonErr != nil {
		return nil, jsonErr
	}
	return &token, nil
}
