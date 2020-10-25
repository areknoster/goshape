package goshape

import (
	"log"

	"fyne.io/fyne/widget"
)



func NewMenu(modes map[string]Mode, setMode func(mode Mode)) *widget.Radio{
	options := make([]string,0, len(modes))
	for modeName := range modes {
		options = append(options, modeName)
	}

	changed := func(modeName string){
		log.Printf("menu set to mode: %s", modeName)
		setMode(modes[modeName])
	}

	radio := widget.NewRadio(options, changed)
	radio.Required = true
	radio.Selected = options[0]
	changed(options[0])

	return radio
}

