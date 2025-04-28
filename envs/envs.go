package envs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	// Carga el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error cargando archivo .env")
	}
}

func Get(env string) string {
	return os.Getenv(env)
}
