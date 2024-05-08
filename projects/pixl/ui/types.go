package ui

import (
	"chaidev.net/pixl/apptype"
	"chaidev.net/pixl/swatch"
	"fyne.io/fyne/v2"
)

type AppInit struct {
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
