package geom

import "math"

type Point struct {
	X, Y float64
}

type Vector struct {
	X, Y float64
}

func (v Vector) SqNorm() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.SqNorm())
}

func MoveByVec(pp, v Pixel) Pixel {
	return Pixel{pp.X + v.X, pp.Y + v.Y}
}