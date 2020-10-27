package goshape

import (
	"goshape/pkg/geom"
)

type ShapesProvider interface {
	GetShapes() (geom.ShapeSet, []RelationsManager)
	SetShapes(geom.ShapeSet, []RelationsManager)

	SelectShape(int)
	GetSelectIndex() int
	UnselectShape()

	SetMode(Mode)
}

type ShapeProvider interface {
	SetShape(geom.Shape, RelationsManager)
	GetSelected() (geom.Shape, RelationsManager)

}

type PlaneProvider interface {
	ShapeProvider
	ShapesProvider
}
