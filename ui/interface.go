package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"recordari/domain"
)

type FyneUI struct {
	window fyne.Window
	a      fyne.App
	list   *widget.List
}

func (i *FyneUI) Mount() *FyneUI {
	a := app.New()
	w := a.NewWindow("Recordari")
	w.Resize(fyne.NewSize(600, 600))
	i.a = a
	i.window = w
	return i
}

func (i *FyneUI) Close() {
	i.window.Close()
}

func (i *FyneUI) ListProjects(userProjects []domain.Project) *FyneUI {
	c := container.NewVBox()
	for _, project := range userProjects {
		card := widget.NewCard(project.Name, "", widget.NewLabel(project.Name))
		border := canvas.NewRectangle(color.White)
		border.StrokeColor = color.Black
		border.StrokeWidth = 1
		cardWithBorder := container.NewStack(
			border,
			container.NewPadded(card),
		)
		c.Add(cardWithBorder)
	}
	i.window.SetContent(c)
	return i
}

func (i *FyneUI) Style(list *widget.List) {

}

func (i *FyneUI) Show() {
	i.window.ShowAndRun()
}
