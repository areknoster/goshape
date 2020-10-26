package goshape

import (
	"goshape/pkg/geom"
)

type ShapesProvider interface {
	GetShapes() geom.ShapeSet
	SetShapes(geom.ShapeSet)
	SelectShape(int)
	GetSelectIndex() int
	SetMode(Mode)
	UnselectShape()
}

type ShapeEditor interface {
	SetShape(geom.Shape)
	GetSelected() geom.Shape
}

type PlaneProvider interface {
	ShapeEditor
	ShapesProvider
}
