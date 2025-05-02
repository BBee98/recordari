package ui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"recordari/ui/settings"
)

func MountWindow() {
	a := app.New()
	w := a.NewWindow("Recordari")
	settings.SetBasicTheme(a)
	settings.SetWindowSize(w)

	w.SetContent(widget.NewLabel("¡Hola Fyne!"))
	w.ShowAndRun()

}
