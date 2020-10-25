package geom

import "math"

type Shape struct{
	Vertices []Vertex
}

func CenteredEquiTriangle(center Point, edgeLen float64) Shape{
	h := math.Sqrt(3) * edgeLen / 2
	v1 := MoveByVector(Vertex(center), Vector{-edgeLen/2, -h/3})
	v2 := MoveByVector(Vertex(center), Vector{edgeLen/2, -h/3})
	v3 := MoveByVector(Vertex(center), Vector{0.0, 2*h/3})
	return Shape{[]Vertex{v1,v2,v3}}
}