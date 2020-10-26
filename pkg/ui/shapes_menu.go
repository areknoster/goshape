package ui

import (
	"log"

	"fyne.io/fyne/widget"

	"goshape/pkg/goshape"
)

func NewShapesMenu(modes []goshape.Mode, set func(mode goshape.Mode)) *widget.Radio {
	options := make([]string, 0, len(modes))
	for _, mode := range modes {
		options = append(options, mode.Name())
	}

	changed := func(modeName string) {
		log.Printf("menu set to mode: %s", modeName)
		for _, mode := range modes {
			if modeName == mode.Name() {
				set(mode)
			}
		}
	}

	radio := widget.NewRadio(options, changed)

	radio.Required = false
	radio.Selected = options[0]

	return radio
}

