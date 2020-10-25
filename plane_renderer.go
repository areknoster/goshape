package goshape

import (
	"image"
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"github.com/sirupsen/logrus"

	"goshape/pkg/geom"
)

type Theme struct {
	backgroundColor color.RGBA
	lineColor       color.RGBA
}

type PlaneRenderer struct {
	render *canvas.Raster
	theme  Theme
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
	return p.theme.backgroundColor
}

func (p *PlaneRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{p.render}
}

func (p *PlaneRenderer) Destroy() {
}

func (p *PlaneRenderer) draw(w, h int) image.Image {
	logrus.Debugf("drawing plane w,h: %v, %v ", w, h)
	rgba := image.NewRGBA(
		image.Rectangle{
			Min: image.Point{0, 0},
			Max: image.Point{w, h},
		})

	putPixel := func(pp geom.PixelPoint){
		rgba.SetRGBA(pp.X, pp.Y, p.theme.lineColor)
	}
	for _, shape := range p.plane.shapes {
		vs := shape.Vertices
		p1 := geom.DenormPoint(geom.Point(vs[len(vs) - 1]), w,h)
		p2 := geom.DenormPoint(geom.Point(vs[0]), w,h)
		Bresenham(p1,p2, putPixel) //between last and first
		for i := 1; i < len(vs); i++ {
			p1 = p2
			p2 = geom.DenormPoint(geom.Point(vs[i]), w,h)
			Bresenham(p1,p2, putPixel)
		}
	}
	return rgba
}

func NewPlaneRenderer(plane *Plane) *PlaneRenderer {
	logrus.Debugf("new plane renderer")
	pr := &PlaneRenderer{
		theme: Theme{
			backgroundColor: colorToRGBA(color.White),
			lineColor:       colorToRGBA(color.Black),
		},
		plane: plane,
	}
	pr.render = canvas.NewRaster(pr.draw)
	return pr
}
