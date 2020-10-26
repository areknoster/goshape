package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/sirupsen/logrus"

	"goshape/pkg/geom"
	"goshape/pkg/goshape"
	"goshape/pkg/modes"
	"goshape/pkg/ui"
)

//	Plane is a widget which contains the logic for the goshape applcation
//	HandleSelect must be set after Plane is initialized
type Plane struct {
	widget.BaseWidget
	mode          goshape.Mode
	size          fyne.Size
	shapes        geom.ShapeSet
	selectedIndex int
	HandleSelect  ui.SetActive
}

func (p *Plane) SetShapes(set geom.ShapeSet) {
	p.shapes = set
}

func (p *Plane) GetSelectIndex() int {
	return p.selectedIndex
}

var _ fyne.Widget = &Plane{}
var _ goshape.ShapesProvider = &Plane{}
var _ goshape.ShapeEditor = &Plane{}

func NewPlane(size fyne.Size) *Plane {
	p := &Plane{
		shapes:        []geom.Shape{},
		size:          size,
		selectedIndex: -1,
	}
	p.mode = modes.NewCreateTriangle(p)
	p.ExtendBaseWidget(p)
	return p
}

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
	normPt := geom.NormPoint(
		geom.Pixel{X: event.Position.X, Y: event.Position.Y},
		p.size.Width, p.size.Height)
	p.mode.HandleClick(normPt)
	logrus.Debugf("Tapped: %v", normPt)
	p.Refresh()
}

func (p *Plane) Dragged(event *fyne.DragEvent) {
	start := geom.NormPoint(
		geom.Pixel{
			X: event.Position.X-event.DraggedX,
			Y: event.Position.Y-event.DraggedY,
		},
		p.size.Width,
		p.size.Height)

	vec := geom.NormVector(
		geom.Pixel{X: event.DraggedX, Y: event.DraggedY},
		p.size.Width, p.size.Height)
	p.mode.HandleDrag(start, vec)
	logrus.Debugf("Drag: start: %v, vec: %v", start, vec)
	p.Refresh()
}

func (p *Plane) DragEnd() {
	logrus.Debugf("Drag finished")
	p.mode.HandleDragEnd()
	p.Refresh()
}

func (p *Plane) SetMode(mode goshape.Mode) {
	p.mode = mode
}

func (p *Plane) GetShapes() geom.ShapeSet {
	return p.shapes
}

func (p *Plane) SetShape(set geom.Shape) {
	p.shapes[p.selectedIndex] = set
}

func (p *Plane) SelectShape(index int) {
	if index < 0 || index >= len(p.shapes) {
		return
	}
	p.selectedIndex = index
	p.HandleSelect(true)
}

func (p *Plane) UnselectShape() {
	p.selectedIndex = -1
	p.HandleSelect(false)
}

func (p *Plane) GetSelected() geom.Shape {
	return p.shapes[p.selectedIndex]
}
