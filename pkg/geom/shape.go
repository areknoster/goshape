package geom

import (
	"math"
)

type Shape struct {
	Vertices []Point
}
func (s Shape) Roll(index int) int{
	return index % len(s.Vertices)
}

type ShapeSet []Shape

type SegmentSet []Segment


func CenteredEquiTriangle(center Point, edgeLen float64) Shape {
	h := math.Sqrt(3) * edgeLen / 2
	v1 := center.MoveByVector(Vector{-edgeLen / 2, -h / 3})
	v2 := center.MoveByVector(Vector{edgeLen / 2, -h / 3})
	v3 := center.MoveByVector(Vector{0.0, 2 * h / 3})
	return Shape{[]Point{v1, v2, v3}}
}

func Middle(s Shape) Point {
	p := Point{0, 0}
	for _, vertex := range s.Vertices {
		p.X, p.Y = p.X+vertex.X, p.Y+vertex.Y
	}
	p.X, p.Y = p.X/float64(len(s.Vertices)), p.Y/float64(len(s.Vertices))
	return p
}

//ClosestToShape returns index of closest shape
func (ss ShapeSet) ClosestToShape(p Point) int {
	if len(ss) == 0 {
		return -1
	}
	minNorm := math.MaxFloat64
	minIndex := -1
	for i, shape := range ss {
		norm := VecBetweenPoints(Middle(shape), p).SqNorm()
		if minNorm > norm {
			minIndex = i
			minNorm = norm
		}
	}
	return minIndex
}

func (s Shape) ClosestToVertex(p Point) int {
	index := -1
	closestSqDist := math.MaxFloat64
	for i, vertex := range s.Vertices {
		if dist := SqDistBetweenPoints(vertex, p); dist < closestSqDist {
			closestSqDist = dist
			index = i
		}
	}
	return index
}

func (s Shape) ToSegmentSet() SegmentSet{
	ss := make([]Segment, len(s.Vertices))
	for i:= 0; i < len(ss) -1; i++ {
		ss[i] = NewSegment(s.Vertices[i], s.Vertices[i+1])
	}
	ss[len(ss) - 1] = NewSegment(s.Vertices[len(s.Vertices) -1], s.Vertices[0])
	return ss
}

func (ss SegmentSet) ShapeFromMidpoints() Shape{
	vs := make([]Point, len(ss))
	for i, s := range ss {
		vs[i] = s.MiddlePoint()
	}
	return Shape{vs}
}

func (s Segment) MoveByVector(v Vector) Segment{
	return Segment{s.A.MoveByVector(v),s.B.MoveByVector(v) }
}
