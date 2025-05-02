package file

import (
	"log"
	"os"
)

func Read(fileName string, msgError string) []byte {
	b, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error al leer el fichero", fileName+" \nMensaje de error :"+msgError)
	}
	return b
}

func CreateAndWriteFile(fileName string, data []byte) *os.File {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Error al crear el fichero", err)
	}
	n, writeErr := file.Write(data)
	if n == 0 || writeErr != nil {
		log.Fatal("Error al escribir en el fichero", writeErr)
	}
	return file
}

func Open(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Error al abrir el fichero", err)
	}
	return file
}
