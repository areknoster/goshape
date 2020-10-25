package goshape

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/sirupsen/logrus"

	"goshape/pkg/geom"
)

//Plane is an abstract representation of canvas.
//Abstract 2D space is normalized to floats: [0,1]X[0,1]
type Plane struct {
	widget.BaseWidget
	mode   Mode
	size   fyne.Size
	shapes []geom.Shape
}

func NewPlane(size fyne.Size) *Plane {
	p := &Plane{
		shapes: []geom.Shape{},
		size:   size,
	}
	p.mode = (*CreateTriangle)(p)
	p.ExtendBaseWidget(p)
	return p
}

var _ fyne.Widget = &Plane{}

func (p *Plane) Size() fyne.Size {
	return p.size
}

func (p *Plane) MinSize() fyne.Size {
	return p.size
}

func (p *Plane) CreateRenderer() fyne.WidgetRenderer {
	return NewPlaneRenderer(p)
}

var _ fyne.Tappable = &Plane{}
var _ fyne.Draggable = &Plane{}

func (p *Plane) Tapped(event *fyne.PointEvent) {
	logrus.Debugf("Tapped: %v", event)
	p.mode.HandleClick(geom.NormPoint(event.Position.X, event.Position.Y, p.size.Width, p.size.Height))
}

func (p *Plane) Dragged(event *fyne.DragEvent) {
	finish := geom.NormPoint(event.Position.X, event.Position.Y, p.size.Width, p.size.Height)
	start := geom.NormPoint(
		event.Position.X-event.DraggedX,
		event.Position.Y-event.DraggedY,
		p.size.Width,
		p.size.Height)
	p.mode.HandleDrag(start, finish)
	p.Refresh()
}

func (p *Plane) DragEnd() {
	p.mode.HandleDragEnd()
	p.Refresh()
}

func (p *Plane) SetMode(mode Mode) {
	p.mode = mode
}
