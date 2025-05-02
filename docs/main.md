# El tipo ``nil``

``nil`` es un tipo exclusivo de `go` que se utiliza para comparar, normalmente, errores.
Esto indica si el error *viene vacío o no*.

#### Ejemplo

````go

package main

import (
	"log"
	"net/http"
)

func main() *http.Response {
	response, err := http.Get("http://this-is-a-fictional-url");
	if err != nil {
		log.Fatal("There was an error doing the request")
	}
	
	return response
}

````

En este ejemplo, la petición ``http.Get("http://this-is-a-fictional-url");`` devuelve dos valores:
- El ``response`` de la petición (es decir, lo que contiene toda la información de la misma)
- El ``error``. Éste es el que debemos averiguar si viene con o sin error.

Esta variable `err` indica cosas como, por ejemplo: 
- La URL no está bien formada
- No hay conexión a internet
Y otro tipo de errores.

Por otro lado, response puede devolver sus propios errores, como un status de `500` que indique que el servicio no está disponible. Por tanto, debemos comprobar
ambos errores, pero la manera de hacerlo es distinta: con el `response` es a partir del código, y con `err` es a través del tipo `nil`.

# Conversión de httpResponses a datos comprensibles para Go



Citando a ChatGPT:

```
Para entenderlo mejor, ``resp.Body`` es como un "tubo" por el que fluyen los datos. Necesitamos una forma de leer esos datos, y ahí es donde entra `io.ReadAll`: `resp.Body`
```


## ¿Qué es ``defer``?

``defer`` es una instrucción que le indica a ``Go`` que **haga algo _cuando_ la función termine**.

Entonces, en esta función:

````go
func ReadFile() ([]byte, error) {
	const credentialsFile = "credentials.json"

	file, fileErr := os.Open(credentialsFile)
	byteFile, byteFileErr := io.ReadAll(file)

	if fileErr != nil || byteFileErr != nil {
		return nil, errors.New("❌ error al leer el archivo")
	}

	fmt.Println(byteFile, "✅ Archivo de credenciales abierto correctamente.")

	defer file.Close()

	return byteFile, nil

}
````

# Aprovecharse de las ventajas de ``Go``como lenguaje idiomático.

Si vienes de front (como es mi caso), es posible que te tiente hacer estructuras
como esta:

````go
type Message struct {
	msg string
	error error
}

func MyFunction(msg) *Message {

	return &Message{msg, nil}
}
````

Algo parecido a como sería en ``typescript``:

````typescript

type Message = {
    msg: string,
    error: Error | undefined
}

function MyFunction(): Message {
    return {
        msg: "Hello world", 
        error: undefined
    }
}
````

Pero esta filosofía **no es del todo correcta para `Go`**.

````
    Incluso la documentación oficial trata el tema: https://go.dev/blog/error-handling-and-go
````


# io.ReadAll vs os.ReadFile

Diferencias:

```
- `io.ReadAll` trabaja con cualquier `io.Reader`
- `os.ReadFile` trabaja específicamente con archivos del sistem
```

Casos de uso:

````
- Usa `os.ReadFile` cuando solo necesites leer un archivo completo del sistema de archivos
- Usa `io.ReadAll` cuando estés trabajando con interfaces `io.Reader` o necesites más flexibilidad en la fuente de datos
````



