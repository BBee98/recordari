package settings

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type basicTheme struct{}

func (m basicTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return &color.RGBA{R: 28, G: 32, B: 48, A: 255} // Fondo principal
	case theme.ColorNameForeground:
		return &color.RGBA{R: 255, G: 255, B: 255, A: 255} // Color del texto
	case theme.ColorNamePrimary:
		return &color.RGBA{R: 65, G: 132, B: 242, A: 255} // Color de énfasis (botones principales)
	case theme.ColorNameButton:
		return &color.RGBA{R: 45, G: 50, B: 65, A: 255} // Color de botones
	case theme.ColorNameHover:
		return &color.RGBA{R: 55, G: 60, B: 75, A: 255} // Color al pasar el mouse
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (m basicTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m basicTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func (m basicTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func SetBasicTheme(app fyne.App) {
	app.Settings().SetTheme(&basicTheme{})
}
