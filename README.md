# recordari

## ¿Qué es ``defer``?

``defer`` es una instrucción que le indica a ``Go`` que **haga algo _cuando_ la función termine**.

Entonces, en esta función:

````
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

La instrucción ``defer file.Close()`` le dice a ``Go`` (en esa función) que cierre el fichero **cuando haga el return**.

## Reglas en Go

### Declaración de variables

- Usa `:=` para declarar variables nuevas

- Usa `=` para asignar a variables ya existentes

### El tipo ``nil``

- ``nil`` indica que **no hubo error**. Así que, muchas veces, cuando realizamos tareas que **pueden ocasionar un error**, devolvemos
``nil`` para indicar que **no lo hubo**.


