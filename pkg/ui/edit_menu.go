package ui

import (
	"log"

	"fyne.io/fyne/widget"

	"goshape/pkg/goshape"
)

type SetActive func(bool)

func NewEditMenu(modes []goshape.Mode, setMode func(mode goshape.Mode)) (*widget.Radio, SetActive) {
	options := make([]string, 0, len(modes))
	for _, mode := range modes {
		options = append(options, mode.Name())
	}

	changed := func(modeName string) {
		log.Printf("menu set to mode: %s", modeName)
		for _, mode := range modes {
			if modeName == mode.Name() {
				setMode(mode)
			}
		}
	}

	radio := widget.NewRadio(options, changed)
	radio.Disable()
	sa := func(b bool) {
		if b {
			radio.Enable()
		} else {
			radio.Disable()
		}

	}
	return radio, sa

}
