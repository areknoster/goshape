package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"

	"goshape/pkg/goshape"
	"goshape/pkg/modes"
)

type Menu struct {
	editMenu   *widget.Radio
	shapesMenu *widget.Radio
}

func NewMenu(sc goshape.PlaneProvider) (*fyne.Container, SetActive) {
	menu := &Menu{}
	cleanSelect := func(mode goshape.Mode) {
		sc.SetMode(mode)
		if menu.editMenu.Selected != mode.Name() {
			menu.editMenu.SetSelected("")
		}
		if menu.shapesMenu.Selected != mode.Name() {
			menu.shapesMenu.SetSelected("")
		}
	}
	editMenu, slider, setActive := NewEditMenu(sc, cleanSelect)
	shapesMenu := NewShapesMenu(modes.NewShapesModesList(sc), cleanSelect)
	menu.editMenu = editMenu
	menu.shapesMenu = shapesMenu

	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(), shapesMenu, editMenu, slider), setActive
}
