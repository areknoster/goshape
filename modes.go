package goshape

import (
	"log"

	"github.com/sirupsen/logrus"

	"goshape/pkg/geom"
)

type Mode interface{
	HandleClick(normLoc geom.Point)
	HandleDrag(start geom.Point, finish geom.Point)
	HandleDragEnd()
}

func NewModesMap(plane *Plane)map[string]Mode {
	return map[string]Mode{
		"Create Triangle" : (*CreateTriangle)(plane),
		"Noop" : (*Noop)(plane),
	}

}


type CreateTriangle Plane

var _ Mode = &CreateTriangle{}

const defaultTriangleEdge = 0.3
func (c *CreateTriangle) HandleClick(normLoc geom.Point) {
	logrus.Debugf("Crate triangle click: %v", normLoc)
	c.shapes = append(c.shapes, geom.CenteredEquiTriangle(normLoc, defaultTriangleEdge))
}

func (c *CreateTriangle) HandleDrag(start geom.Point, finish geom.Point) {
	// no operation on drag
}

func (c *CreateTriangle) HandleDragEnd() {
	// no operation on drag

}

type Noop Plane
var _ Mode = &Noop{}


func (n Noop) HandleClick(normLoc geom.Point) {
	log.Printf("click: %v", normLoc)
}

func (n Noop) HandleDrag(start geom.Point, finish geom.Point) {
	log.Printf("drag: %v", start)

}

func (n *Noop) HandleDragEnd() {
	log.Printf("drag end")


}





