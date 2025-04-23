# FYNE

## Instalación

Para poder empezar a utilizar ``fyne``, debemos instalar el paquete:

```
go get fyne.io/fyne/v2@latest
```

Y, a continuación, importarlo para utilizarlo:

## Empezar a usar Fyne. Ejemplo básico y sencillo

Vamos a mirar esta función de ejemplo:

```
package pop_up

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func Show() {
	a := app.New()
	window := a.NewWindow("Hello World")

	window.SetContent(widget.NewLabel("Hello World!"))
	window.ShowAndRun()
}


```

Toda aplicación realizada con ``fyne`` debe empezar con la misma instrucción:

```
	a := app.New()
```

Esta instrucción indica que **estamos creando una nueva aplicación de ``fyne```**, y nos facilita el acceso a una serie de herramientas 
para realizar nuestra aplicación de escritorio.

El resto es sencillo: 

````
    window := a.NewWindow("Hello World")

	window.SetContent(widget.NewLabel("Hello World!"))
	window.ShowAndRun()
````

- Creamos una ventana ``a.NewWindow``
- Le decimos a esa ventana que tenga un contenido ``window.SetContent``, el cuál va a ser un simple texto.
- Por último, le decimos ``ShowAndRun`` que se muestre.


Esto es tan solo una muestra, pero lo que queremos hacer es un poco diferente.
