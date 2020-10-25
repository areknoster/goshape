package main

import (
	"fyne.io/fyne"
	fyneapp "fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"github.com/sirupsen/logrus"

	"goshape"
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

	plane := goshape.NewPlane(defaultCfg.boardSize)
	menu := goshape.NewMenu(goshape.NewModesMap(plane), plane.SetMode)

	container := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, menu, nil), menu, plane )
	window.SetContent(container)
	window.SetFixedSize(true)


	window.ShowAndRun()

}
