package settings

import "fyne.io/fyne/v2"

func SetWindowSize(window fyne.Window) {
	window.Resize(fyne.NewSize(600, 600))
}
