package goshape

import "goshape/pkg/geom"

type Mode interface {
	HandleClick(normLoc geom.Point)
	HandleDrag(start geom.Point, move geom.Vector)
	HandleDragEnd()
	Name() string
}

