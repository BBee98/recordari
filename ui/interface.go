package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"recordari/domain"
	"recordari/ui/settings"
)

type FyneUI struct {
	window fyne.Window
	a      fyne.App
	list   *widget.List
}

func (i *FyneUI) Mount() *FyneUI {
	a := app.New()
	w := a.NewWindow("Recordari")
	settings.SetBasicTheme(a)
	settings.SetWindowSize(w)
	i.a = a
	i.window = w
	return i
}

func (i *FyneUI) Close() {
	i.window.Close()
}

func (i *FyneUI) ListProjects(userProjects []domain.Project) *FyneUI {
	list := widget.NewList(
		func() int {
			return len(userProjects)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Proyectos")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(userProjects[i].Name)
		})
	i.window.SetContent(list)

	return i
}

func (i *FyneUI) Style(list *widget.List) {

}

func (i *FyneUI) Show() {
	i.window.ShowAndRun()
}
