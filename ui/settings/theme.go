package settings

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type basicTheme struct{}

func (m basicTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
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
