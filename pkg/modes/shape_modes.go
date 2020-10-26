package modes

import (
	"github.com/sirupsen/logrus"

	"goshape/pkg/geom"
	"goshape/pkg/goshape"
)

func NewShapesModesList(sc goshape.ShapesProvider) []goshape.Mode {
	return []goshape.Mode{
		NewCreateTriangle(sc),
		NewSelectShape(sc),
		NewDeleteShape(sc),
	}

}

type CreateTriangle struct {
	goshape.ShapesProvider
}

func NewCreateTriangle(shapesContainer goshape.ShapesProvider) *CreateTriangle {
	return &CreateTriangle{ShapesProvider: shapesContainer}
}

var _ goshape.Mode = CreateTriangle{}

func (c CreateTriangle) Name() string {
	return "Create Triangle"
}

func (c CreateTriangle) HandleClick(normLoc geom.Point) {
	const defaultTriangleEdge = 0.3
	shapes := c.ShapesProvider.GetShapes()
	shapes = append(shapes, geom.CenteredEquiTriangle(normLoc, defaultTriangleEdge))
	c.SetShapes(shapes)
}

func (c CreateTriangle) HandleDrag(start geom.Point, move geom.Vector) {}

func (c CreateTriangle) HandleDragEnd() {}

type SelectShape struct {
	goshape.ShapesProvider
}

func NewSelectShape(shapesContainer goshape.ShapesProvider) *SelectShape {
	return &SelectShape{ShapesProvider: shapesContainer}
}

var _ goshape.Mode = SelectShape{}

func (s SelectShape) Name() string {
	return "Select Shape"
}

func (s SelectShape) HandleClick(normLoc geom.Point) {
	c := s.GetShapes().ClosestToShape(normLoc)
	if c < 0 {
		s.UnselectShape()
		return
	}
	s.SelectShape(c)
}

func (s SelectShape) HandleDrag(start geom.Point, move geom.Vector) {}

func (s SelectShape) HandleDragEnd() {}

type DeleteShape struct {
	goshape.ShapesProvider
}

func NewDeleteShape(shapesContainer goshape.ShapesProvider) *DeleteShape {
	return &DeleteShape{ShapesProvider: shapesContainer}
}

var _ goshape.Mode = DeleteShape{}

func (d DeleteShape) Name() string {
	return "Delete Shape"
}

func (d DeleteShape) HandleClick(normLoc geom.Point) {
	shapes := d.GetShapes()
	clicked := shapes.ClosestToShape(normLoc)
	switch {
	case clicked == -1:
		return
	case clicked == d.GetSelectIndex():
		d.UnselectShape()
	case clicked < d.GetSelectIndex():
		d.SelectShape(d.GetSelectIndex() - 1)
		logrus.Debugf("deleting shape with index: %d", clicked)
	default:
		logrus.Debugf("deleting shape with index: %d", clicked)
	}
	shapes = append(shapes[0:clicked], shapes[clicked+1:]...)
	d.SetShapes(shapes)
}

func (d DeleteShape) HandleDrag(start geom.Point, move geom.Vector) {}

func (d DeleteShape) HandleDragEnd() {}
