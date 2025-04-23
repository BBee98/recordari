package google_credentials

import (
	"fmt"
	"io"
	"os"
)

func Get() ([]byte, error) {
	const credentialsFile = "credentials.json"

	file, fileErr := os.Open(credentialsFile)
	byteFile, byteFileErr := io.ReadAll(file)

	if fileErr != nil || byteFileErr != nil {
		err := fmt.Errorf("ha habido un error al abrir el fichero de credenciales: %v", fileErr)
		return nil, err
	}

	fmt.Println(byteFile, "✅ Archivo de credenciales abierto correctamente.")

	defer file.Close()

	return byteFile, nil

}
