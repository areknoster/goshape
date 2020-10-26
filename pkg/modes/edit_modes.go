package modes

import (
	"goshape/pkg/geom"
	"goshape/pkg/goshape"
)

func NewEditModesList(se goshape.ShapeEditor) []goshape.Mode {
	return []goshape.Mode{
		NewMoveVertex(se),
		NewAddVertex(se),
		NewMoveSegment(se),
		NewMoveShape(se),
	}
}

type MoveVertex struct {
	goshape.ShapeEditor
	curVert int
}

func NewMoveVertex(ShapeEditor goshape.ShapeEditor) *MoveVertex {
	return &MoveVertex{ShapeEditor: ShapeEditor, curVert: -1}
}

func (m *MoveVertex) HandleClick(normLoc geom.Point) {}

func (m *MoveVertex) HandleDrag(start geom.Point, move geom.Vector) {
	shape := m.GetSelected()
	if m.curVert == -1 {
		m.curVert = shape.ClosestToVertex(start)
	}

	newVert := shape.Vertices[m.curVert].MoveByVector(move)
	shape.Vertices[m.curVert] = newVert
	m.SetShape(shape)
}

func (m *MoveVertex) HandleDragEnd() {
	m.curVert = -1
}

func (m *MoveVertex) Name() string {
	return "Move Vertex"
}

type AddVertex struct {
	goshape.ShapeEditor
}

func NewAddVertex(ShapeEditor goshape.ShapeEditor) *AddVertex {
	return &AddVertex{ShapeEditor: ShapeEditor}
}

func (a AddVertex) HandleClick(normLoc geom.Point) {
	shape := a.GetSelected()
	midPts := shape.ToSegmentSet().ShapeFromMidpoints()
	index := midPts.ClosestToVertex(normLoc)
	vs := shape.Vertices

	switch{
	case index == len(vs) - 1:
		vs = append(vs, midPts.Vertices[index])
	default:
		vtmp := make([]geom.Point, len(vs) + 1)
		copy(vtmp[:index+1], vs[:index+1])
		vtmp[index+1] = midPts.Vertices[index]
		copy(vtmp[index+2:], vs[index+1:])
		vs = vtmp
	}
	shape.Vertices = vs
	a.SetShape(shape)

}

func (a AddVertex) HandleDrag(start geom.Point, move geom.Vector) {}
func (a AddVertex) HandleDragEnd()                                {}

func (a AddVertex) Name() string {
	return "Add Vertex"
}

type MoveSegment struct{
	goshape.ShapeEditor
	currSegment int
}

func NewMoveSegment(ShapeEditor goshape.ShapeEditor) *MoveSegment {
	return &MoveSegment{ShapeEditor: ShapeEditor, currSegment: -1}
}

func (m *MoveSegment) HandleClick(normLoc geom.Point) {}

func (m *MoveSegment) HandleDrag(start geom.Point, move geom.Vector) {
	shape := m.GetSelected()
	vs := shape.Vertices
	segments := shape.ToSegmentSet()
	if m.currSegment == -1 {
		segmentsMids := segments.ShapeFromMidpoints()
		m.currSegment = segmentsMids.ClosestToVertex(start)
	}

	switch{
	case m.currSegment == len(segments) -1:
		vs[m.currSegment] = vs[m.currSegment].MoveByVector(move)
		vs[0] = vs[0].MoveByVector(move)
	default:
		vs[m.currSegment] = vs[m.currSegment].MoveByVector(move)
		vs[m.currSegment + 1] = vs[m.currSegment+1].MoveByVector(move)
	}
	shape.Vertices = vs
	m.SetShape(shape)
}

func (m *MoveSegment) HandleDragEnd() {
	m.currSegment = -1
}

func (m *MoveSegment) Name() string {
	return "Move Segment"
}

type MoveShape struct{
	goshape.ShapeEditor
}

func NewMoveShape(ShapeEditor goshape.ShapeEditor) MoveShape {
	return MoveShape{ShapeEditor: ShapeEditor}
}

func (m MoveShape) HandleClick(normLoc geom.Point) {}

func (m MoveShape) HandleDrag(start geom.Point, move geom.Vector) {
	shape := m.GetSelected()
	vs := shape.Vertices
	for i, v := range vs {
		vs[i] = v.MoveByVector(move)
	}
	shape.Vertices = vs
	m.SetShape(shape)

}

func (m MoveShape) HandleDragEnd() {}

func (m MoveShape) Name() string {
	return "Move Shape"
}






