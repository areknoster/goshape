package geom

import (
	"math"
)

func (v Point) MoveByVector(vec Vector) Point {
	return Point{
		X: v.X + vec.X,
		Y: v.Y + vec.Y,
	}
}

func VecBetweenPoints(v1, v2 Point) Vector {
	return Vector{v2.X - v1.X, v2.Y - v1.Y}
}

func VecLenSq(v Vector) float64 {
	return v.X*v.X + v.Y*v.Y
}

func SqDistBetweenPoints(a, b Point) float64 {
	return VecLenSq(VecBetweenPoints(a, b))
}


func VecNorm(v0, v1 Vector) float64{
	return v0.X * v1.X + v0.Y * v1.Y
}

var ZeroVector = Vector{0,0}
var WrongPoint = Point{math.NaN(), math.NaN()}

func ClosestPoint(P, A, B Point) Point {
	vecAB := VecBetweenPoints(A,B)
	vecAP := VecBetweenPoints(A,P)
	normK := VecNorm(vecAB, vecAP)
	normL := VecNorm(vecAB, vecAB)
	slope := -normK / normL

	f := func(s float64) Point {
		return Point{
			X : (1.0 - slope) * A.X + slope * B.X - P.X,
			Y : (1.0 - slope) * A.Y + slope * B.Y - P.Y,
		}
	}

	if slope < 0 || slope > 1{
		return WrongPoint
	}
	return f(slope)


}

func (v Vector) TimesScalar(s float64) Vector{
	return Vector{s * v.X, s* v.Y}
}