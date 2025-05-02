package credentials

import (
	"fmt"
	"recordari/envs"
	"recordari/services/file"
)

func GetGoogleCredentials(msgError string) []byte {

	byteFile := file.Read(envs.Get("GOOGLE_CREDENTIALS"), msgError)
	fmt.Println("✅ Archivo de credenciales abierto correctamente.")

	return byteFile

}
