# Variables globales

Para declarar variables globales, debe de hacerse fuera de cualquier función.
Dependiendo de qué intenciones tengamos con esas variables, debemos usar una estructura u otra:

1. PascalCase sin "_" para constantes (al igual que en `javscript`, usamos la key `const`)
2. camelCase para variables que queramos usar/compartir con otros ficheros, o para evitar re-declararlas (por ejemplo, para aplicar el patrón `singleton`). Aqui podemos usar
``var`` si necesitamos asignarle un valor más tarde del momento en el que la declaramos, o ``const``

#### Ejemplo

```go
package main

const ConstantVariable = "foo"
var globalVariable any

func main(){
	println(ConstantVariable)
	println(globalVariable)
}
```

# Asignar vs Declarar

- Usa `:=` para declarar variables nuevas

- Usa `=` para asignar a variables ya existentes

#### Ejemplo

```go

package main

type Message struct {
	text string
}

func main() {
	
	newVariableDeclared := Message{
		text: "Hello",
	}
	
	println(newVariableDeclared)
}

```

# Inyección de dependencia

En ``Go`` no hay inyección de dependencia nativa, pero puede hacerse mediante **constructores**

#### Ejemplo:

````go

// Definición de la interfaz
type ProjectsRepository interface {
    GetAll() ([]*Project, error)
}

// Handler con DI
type GetProjectsHandler struct {
    projectsRepository ProjectsRepository
}

// Constructor que implementa la inyección
func NewGetProjectsHandler(repo ProjectsRepository) *GetProjectsHandler {
    return &GetProjectsHandler{
    projectsRepository: repo,
    }
}

// Método que usa la dependencia inyectada
func (h *GetProjectsHandler) Handle() ([]*Project, error) {
    return h.projectsRepository.GetAll()
}
````


# Usa siempre defer Body.Close() en una función donde obtengas el cuerpo de una petición http y dónde lo leas.

La instrucción ``defer file.Close()`` le dice a ``Go`` (en esa función) que cierre el fichero **cuando haga el return**.


