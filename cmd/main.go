package main

import (
	"fyne.io/fyne"
	fyneapp "fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"github.com/sirupsen/logrus"

	"goshape/pkg/ui"
)

type Config struct{
	title string
	boardSize fyne.Size
}

var defaultCfg = Config{
	title: "GoShape",
	boardSize: fyne.NewSize(600, 600),
}

func main(){
	logrus.SetLevel(logrus.DebugLevel)
	cfg := defaultCfg
	fyneApp := fyneapp.New()
	window :=fyneApp.NewWindow(cfg.title)
	plane := NewPlane(defaultCfg.boardSize)

	menu, setActiveEditMenu := ui.NewMenu(plane)
	plane.HandleSelect = setActiveEditMenu

	container := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, menu, nil), menu, plane )
	window.SetContent(container)
	window.SetFixedSize(true)


	window.ShowAndRun()

}
