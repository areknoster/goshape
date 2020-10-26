package geom

type Segment struct{
	A,B Point
}

func NewSegment(A,B Point) Segment {
	return Segment{A,B}
}

func (s Segment) ToVector() Vector{
	return VecBetweenPoints(s.A, s.B)
}

func (s Segment) MiddlePoint() Point{
	return s.A.MoveByVector(s.ToVector().TimesScalar(0.5))
}





