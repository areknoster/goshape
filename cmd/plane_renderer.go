package main

import (
	"image"
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"github.com/sirupsen/logrus"

	"goshape/pkg/goshape"
	"goshape/pkg/render"
)

type PlaneRenderer struct {
	render *canvas.Raster
	theme  render.Theme
	plane  *Plane
}

var _ fyne.WidgetRenderer = &PlaneRenderer{}

func (p *PlaneRenderer) Layout(size fyne.Size) {
	p.render.Resize(size)
}

func (p *PlaneRenderer) MinSize() fyne.Size {
	return p.plane.Size()
}

func (p *PlaneRenderer) Refresh() {
	logrus.Debug("PlaneRenderer: refresh")
	canvas.Refresh(p.render)
}

func (p *PlaneRenderer) BackgroundColor() color.Color {
	return p.theme.BackgroundColor
}

func (p *PlaneRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{p.render}
}

func (p *PlaneRenderer) Destroy() {
}

func (p *PlaneRenderer) draw(w, h int) image.Image {
	logrus.Debugf("drawing plane w,h: %v, %v ", w, h)
	img := image.NewRGBA(
		image.Rectangle{
			Min: image.Point{0, 0},
			Max: image.Point{w, h},
		})

	for i, shape := range p.plane.shapes {
		if i != p.plane.selectedIndex{
			logrus.Debug("rendering common shape")
			render.RenderCommonShape(shape, img, p.theme)
		}else {
			logrus.Debug("rendering selected shape")
			render.RenderSelectedShape(shape, img, p.theme)
		}
	}
	return img
}



func NewPlaneRenderer(plane *Plane) *PlaneRenderer {
	logrus.Debugf("new plane renderer")
	pr := &PlaneRenderer{
		theme: render.Theme{
			BackgroundColor: goshape.ColorToRGBA(color.White),
			LineColor:       goshape.ColorToRGBA(color.Black),
			AccentColor:     color.RGBA{200, 0, 0, 255 },
		},
		plane: plane,
	}
	pr.render = canvas.NewRaster(pr.draw)
	return pr
}
