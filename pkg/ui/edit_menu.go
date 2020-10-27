package ui

import (
	"log"
	"math"

	"fyne.io/fyne/widget"

	"goshape/pkg/goshape"
	modes2 "goshape/pkg/modes"
)

type SetActive func(bool)

func NewEditMenu(sp goshape.ShapeProvider, setMode func(mode goshape.Mode)) (*widget.Radio, *widget.Slider, SetActive) {
	modes, angleSetter := modes2.NewEditModesList(sp)
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

	slider := widget.NewSlider(0.0, math.Pi)
	slider.OnChanged = angleSetter
	return radio, slider, sa

}
