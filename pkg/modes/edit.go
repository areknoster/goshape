package modes

import (
	"goshape/pkg/geom"
	"goshape/pkg/goshape"
	"goshape/pkg/relation"
)



func NewEditModesList(sp goshape.ShapeProvider) ([]goshape.Mode, func(float64)) {
	angle, setter := NewSetAngle(sp)
	return []goshape.Mode{
		NewMoveVertex(sp),
		NewAddVertex(sp),
		NewDeleteVertex(sp),
		NewMoveSegment(sp),
		NewMoveShape(sp),
		NewSetVertical(sp),
		NewSetHorizontal(sp),
		NewDeleteRelation(sp),
		NewMoveAfterReleased(sp),
		angle,
	}, setter
}

type MoveVertex struct {
	goshape.ShapeProvider
	curVert int
}

func NewMoveVertex(ShapeEditor goshape.ShapeProvider) *MoveVertex {
	return &MoveVertex{ShapeProvider: ShapeEditor, curVert: -1}
}

func (m *MoveVertex) HandleClick(normLoc geom.Point) {}

func (m *MoveVertex) HandleDrag(start geom.Point, move geom.Vector) {
	shape, rm := m.GetSelected()
	if m.curVert == -1 {
		m.curVert = shape.ClosestToVertex(start)
	}

	newVert := shape.Vertices[m.curVert].MoveByVector(move)
	shape.Vertices[m.curVert] = newVert
	shape = rm.ApplyRelations(shape, m.curVert)
	m.SetShape(shape, rm)
}

func (m *MoveVertex) HandleDragEnd() {
	m.curVert = -1
}

func (m *MoveVertex) Name() string {
	return "Move Vertex"
}

type AddVertex struct {
	goshape.ShapeProvider
}

func NewAddVertex(ShapeEditor goshape.ShapeProvider) *AddVertex {
	return &AddVertex{ShapeProvider: ShapeEditor}
}

func (a AddVertex) HandleClick(normLoc geom.Point) {
	shape, rm := a.GetSelected()
	midPts := shape.ToSegmentSet().ShapeFromMidpoints()
	index := midPts.ClosestToVertex(normLoc)

	vs := shape.Vertices

	switch {
	case index == len(vs)-1:
		vs = append(vs, midPts.Vertices[index])
	default:
		vtmp := make([]geom.Point, len(vs)+1)
		copy(vtmp[:index+1], vs[:index+1])
		vtmp[index+1] = midPts.Vertices[index]
		copy(vtmp[index+2:], vs[index+1:])
		vs = vtmp
	}
	shape.Vertices = vs

	rm = rm.IncrementOver(index, shape.Roll)
	rm.RemoveCollisions(index)

	a.SetShape(shape, rm)

}

func (a AddVertex) HandleDrag(start geom.Point, move geom.Vector) {}
func (a AddVertex) HandleDragEnd()                                {}

func (a AddVertex) Name() string {
	return "Add Vertex"
}

type DeleteVertex struct {
	goshape.ShapeProvider
}

func (d DeleteVertex) HandleClick(normLoc geom.Point) {
	shape, rm := d.GetSelected()
	vs := shape.Vertices
	if len(shape.Vertices) == 3 {
		return
	}
	v := shape.ClosestToVertex(normLoc)
	rm.RemoveCollisions(v)
	vs = append(vs[:v], vs[v+1:]...)
	shape.Vertices = vs
	d.SetShape(shape, rm)
}

func (d DeleteVertex) HandleDrag(start geom.Point, move geom.Vector) {}

func (d DeleteVertex) HandleDragEnd() {}

func (d DeleteVertex) Name() string {
	return "Delete Vertex"
}

func NewDeleteVertex(shapeProvider goshape.ShapeProvider) DeleteVertex {
	return DeleteVertex{ShapeProvider: shapeProvider}
}

type MoveSegment struct {
	goshape.ShapeProvider
	currSegment int
}

func NewMoveSegment(ShapeEditor goshape.ShapeProvider) *MoveSegment {
	return &MoveSegment{ShapeProvider: ShapeEditor, currSegment: -1}
}

func (m *MoveSegment) HandleClick(normLoc geom.Point) {}

func (m *MoveSegment) HandleDrag(start geom.Point, move geom.Vector) {
	shape, rm := m.GetSelected()
	vs := shape.Vertices
	segments := shape.ToSegmentSet()
	if m.currSegment == -1 {
		segmentsMids := segments.ShapeFromMidpoints()
		m.currSegment = segmentsMids.ClosestToVertex(start)
	}
	switch {
	case m.currSegment == len(segments)-1:
		vs[m.currSegment] = vs[m.currSegment].MoveByVector(move)
		vs[0] = vs[0].MoveByVector(move)
	default:
		vs[m.currSegment] = vs[m.currSegment].MoveByVector(move)
		vs[m.currSegment+1] = vs[m.currSegment+1].MoveByVector(move)
	}
	shape.Vertices = vs
	rm.ApplyRelations(shape, m.currSegment, m.currSegment+1)
	m.SetShape(shape, rm)
}

func (m *MoveSegment) HandleDragEnd() {
	m.currSegment = -1
}

func (m *MoveSegment) Name() string {
	return "Move Segment"
}

type MoveShape struct {
	goshape.ShapeProvider
}

func NewMoveShape(ShapeEditor goshape.ShapeProvider) MoveShape {
	return MoveShape{ShapeProvider: ShapeEditor}
}

func (m MoveShape) HandleClick(normLoc geom.Point) {}

func (m MoveShape) HandleDrag(start geom.Point, move geom.Vector) {
	shape, rm := m.GetSelected()
	vs := shape.Vertices
	for i, v := range vs {
		vs[i] = v.MoveByVector(move)
	}
	shape.Vertices = vs
	m.SetShape(shape, rm)

}

func (m MoveShape) HandleDragEnd() {}

func (m MoveShape) Name() string {
	return "Move Shape"
}

type SetVertical struct {
	goshape.ShapeProvider
}

func (sv SetVertical) HandleClick(normLoc geom.Point) {
	shape, rm := sv.GetSelected()
	vertice := shape.ToSegmentSet().ShapeFromMidpoints().ClosestToVertex(normLoc)
	rel := relation.NewVerticalSegment(vertice, vertice+1, shape.Roll)
	rm = rm.RemoveCollisions(vertice, vertice + 1)
	rm = rm.AddRelation(rel)
	sv.SetShape(shape, rm)
}

func (sv SetVertical) HandleDrag(start geom.Point, move geom.Vector) {}

func (sv SetVertical) HandleDragEnd() {}

func (sv SetVertical) Name() string {
	return "Set Vertical"
}

func NewSetVertical(shapeProvider goshape.ShapeProvider) SetVertical {
	return SetVertical{ShapeProvider: shapeProvider}
}

type SetHorizontal struct {
	goshape.ShapeProvider
}

func (sh SetHorizontal) HandleClick(normLoc geom.Point) {
	shape, rm := sh.GetSelected()
	vertice := shape.ToSegmentSet().ShapeFromMidpoints().ClosestToVertex(normLoc)
	rel := relation.NewHorizontalSegment(vertice, vertice+1, shape.Roll)
	rm = rm.RemoveCollisions(vertice, vertice + 1)
	rm = rm.AddRelation(rel)
	sh.SetShape(shape, rm)
}

func (sh SetHorizontal) HandleDrag(start geom.Point, move geom.Vector) {}

func (sh SetHorizontal) HandleDragEnd() {}

func (sh SetHorizontal) Name() string {
	return "Set Horizontal"
}

func NewSetHorizontal(shapeProvider goshape.ShapeProvider) SetHorizontal {
	return SetHorizontal{ShapeProvider: shapeProvider}
}

type DeleteRelation struct {
	goshape.ShapeProvider
}

func NewDeleteRelation(shapeProvider goshape.ShapeProvider) DeleteRelation {
	return DeleteRelation{ShapeProvider: shapeProvider}
}

var _ goshape.Mode = DeleteRelation{}

func (d DeleteRelation) HandleClick(normLoc geom.Point) {
	sh, rm := d.GetSelected()
	rm = rm.DeleteRelation(sh, normLoc)
	d.SetShape(sh, rm)
}

func (d DeleteRelation) HandleDrag(start geom.Point, move geom.Vector) {}

func (d DeleteRelation) HandleDragEnd() {}

func (d DeleteRelation) Name() string {
	return "Delete Relation"
}

type SetAngle struct{
	goshape.ShapeProvider
	angle float64
}

func (s SetAngle) HandleClick(normLoc geom.Point) {
}

func (s SetAngle) HandleDrag(start geom.Point, move geom.Vector) {
}

func (s SetAngle) HandleDragEnd() {
}

func (s SetAngle) Name() string {
	return "Set Angle"
}

func NewSetAngle(shapeProvider goshape.ShapeProvider) (*SetAngle, func(float64)){
	sa :=  &SetAngle{ShapeProvider: shapeProvider}
	setter := func(angle float64){
		sa.angle = angle
	}
	return &SetAngle{ShapeProvider: shapeProvider}, setter
}

//Zadanie domowe

type MoveAfterReleased struct {
	goshape.ShapeProvider
	curVert int
	lastRes geom.Point
}

func NewMoveAfterReleased(shapeProvider goshape.ShapeProvider) *MoveAfterReleased {
	return &MoveAfterReleased{ShapeProvider: shapeProvider, curVert: -1}
}

func (m *MoveAfterReleased) HandleClick(normLoc geom.Point) {
}

func (m *MoveAfterReleased) HandleDrag(start geom.Point, move geom.Vector) {
	sh, rm := m.GetSelected()
	if m.curVert == -1{
		m.curVert = sh.ClosestToVertex(start)
	}
	m.lastRes = start.MoveByVector(move)
	m.SetShape(sh, rm)
}

func (m *MoveAfterReleased) HandleDragEnd() {
	sh, rm := m.GetSelected()
	vs := sh.Vertices
	vs[m.curVert] = m.lastRes
	sh.Vertices = vs
	m.SetShape(sh, rm)

}

func (m *MoveAfterReleased) Name() string {
	return "LabB: LeaveShadow"
}




