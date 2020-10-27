package relation

import (
	"errors"
	"image"
	"image/color"
	"math"

	"github.com/sirupsen/logrus"

	"goshape/pkg/geom"
	"goshape/pkg/goshape"
	"goshape/pkg/render"
)

type VerticalSegment struct {
	points []int
	roll   goshape.Roll
}

var _ goshape.Relation = VerticalSegment{}

func (v VerticalSegment) IncrementOver(index int, roll goshape.Roll) goshape.Relation {
	np := v.points
	for i, point := range v.points {
		if point > index {
			np[i] = roll(np[i] + 1)
		}
	}
	return VerticalSegment{
		points: np,
		roll:   roll,
	}
}

func (v VerticalSegment) GetPoints() []int {
	return v.points
}

func (v VerticalSegment) Repair(shape geom.Shape, fixedPoints []int) geom.Shape {
	logrus.Debug("repairing shape")
	vs := shape.Vertices
	for _, point := range fixedPoints {
		switch {
		case point == v.points[0]:
			vs[shape.Roll(v.points[1])].X = vs[shape.Roll(v.points[0])].X
			shape.Vertices = vs
		case point == v.points[1]:
			vs[shape.Roll(v.points[0])].X = vs[shape.Roll(v.points[1])].X
			shape.Vertices = vs
		}
	}
	vs[shape.Roll(v.points[0])].X = vs[shape.Roll(v.points[1])].X
	return shape

}

func (v VerticalSegment) Draw(img *image.RGBA, shape geom.Shape) {
	w, h := img.Rect.Max.X, img.Rect.Max.Y
	vs := shape.Vertices
	mid := geom.NewSegment(vs[shape.Roll(v.points[0])], vs[shape.Roll(v.points[1])]).MiddlePoint()
	const step = 0.02
	lowV := mid.MoveByVector(geom.Vector{step, -step})
	highV := lowV.MoveByVector(geom.Vector{0, 2 * step})
	low := geom.DenormPoint(lowV, w, h)
	high := geom.DenormPoint(highV, w, h)
	brush := render.NewSquareBrush(2, img, color.RGBA{0, 0, 0, 255})

	render.BresenhamLine(low, high, brush)

}

func (v VerticalSegment) GetMiddle(shape geom.Shape) geom.Point {
	vs := shape.Vertices
	seg := geom.Segment{vs[v.points[0]], vs[v.points[1]]}
	return seg.MiddlePoint()
}

func (v VerticalSegment) Collide(pts ...int) bool {
	for _, pt := range pts {
		if pt == v.points[0] || pt == v.points[1] {
			return true
		}
	}
	return false
}

func NewVerticalSegment(a, b int, roll goshape.Roll) *VerticalSegment {
	return &VerticalSegment{
		points: []int{a, b},
		roll:   roll,
	}
}

type HorizontalSegment struct {
	points []int
	roll   goshape.Roll
}

var _ goshape.Relation = HorizontalSegment{}

func (hs HorizontalSegment) IncrementOver(index int, roll goshape.Roll) goshape.Relation {
	np := hs.points
	for i, point := range hs.points {
		if point > index {
			np[i] = roll(np[i] + 1)
		}
	}
	return HorizontalSegment{
		points: np,
		roll:   roll,
	}
}

func (hs HorizontalSegment) GetPoints() []int {
	return hs.points
}

func (hs HorizontalSegment) Repair(shape geom.Shape, fixedPoints []int) geom.Shape {
	logrus.Debug("repairing shape")
	vs := shape.Vertices
	for _, point := range fixedPoints {
		switch {
		case point == hs.points[0]:
			vs[shape.Roll(hs.points[1])].Y = vs[shape.Roll(hs.points[0])].Y
			shape.Vertices = vs
		case point == hs.points[1]:
			vs[shape.Roll(hs.points[0])].Y = vs[shape.Roll(hs.points[1])].Y
			shape.Vertices = vs
		}
	}
	vs[shape.Roll(hs.points[0])].Y = vs[shape.Roll(hs.points[1])].Y
	return shape

}

func (hs HorizontalSegment) Draw(img *image.RGBA, shape geom.Shape) {
	w, h := img.Rect.Max.X, img.Rect.Max.Y
	vs := shape.Vertices
	mid := geom.NewSegment(vs[shape.Roll(hs.points[0])], vs[shape.Roll(hs.points[1])]).MiddlePoint()
	const step = 0.02
	lowV := mid.MoveByVector(geom.Vector{-step, -step})
	highV := lowV.MoveByVector(geom.Vector{2 * step, 0})
	low := geom.DenormPoint(lowV, w, h)
	high := geom.DenormPoint(highV, w, h)
	brush := render.NewSquareBrush(2, img, color.RGBA{0, 0, 0, 255})

	render.BresenhamLine(low, high, brush)

}

func (hs HorizontalSegment) GetMiddle(shape geom.Shape) geom.Point {
	vs := shape.Vertices
	seg := geom.Segment{vs[hs.points[0]], vs[hs.points[1]]}
	return seg.MiddlePoint()
}

func (hs HorizontalSegment) Collide(pts ...int) bool {
	for _, pt := range pts {
		if pt == hs.roll(hs.points[0]) || pt == hs.roll(hs.points[1]) {
			return true
		}
	}
	return false
}

func NewHorizontalSegment(a, b int, roll goshape.Roll) HorizontalSegment {
	return HorizontalSegment{
		points: []int{a, b},
		roll:   roll,
	}
}

type Angle struct {
	a, b, c int
	roll    goshape.Roll
	value   float64
}

var _ goshape.Relation = Angle{}

func isNextTo(a, b int, roll goshape.Roll) bool {
	if roll(a)+1 == b || roll(b)+1 == a {
		return true
	}
	return false
}

func NewAngle(a, b, c int, roll goshape.Roll, angle float64) (Angle, error) {
	if !isNextTo(a, b, roll) || !isNextTo(b, c, roll) {
		return Angle{}, errors.New("points are not tex to each other skipping")

	}
	return Angle{
		a: a, b: b, c: c,
		roll: roll, value: angle,
	}, nil
}

func (a Angle) Repair(shape geom.Shape, fixedPoints []int) geom.Shape {
	isIn := func(p int) bool {
		for _, f := range fixedPoints {
			if f == p {
				return true
			}
		}
		return false
	}

	yOfv2 := func(v1, v2 geom.Vector) float64 {
		n := geom.VecNorm(v1, v2)
		nSq := n * n
		ang := math.Cos(a.value)
		angSq := ang * ang
		return math.Sqrt(nSq/(angSq*v1.SqNorm()) - v2.X*v2.X)
	}
	ia, ib, ic := isIn(a.a), isIn(a.b), isIn(a.c)

	vs := shape.Vertices

	switch {
	case !ia && !ib && !ic:
		logrus.Debug("angle: nothing to repare")
	case ia && ib && ic:
		panic("three points cannot be fixed!")
	case ia && ic:
		panic("both first and last cannot be fixed")
	case ia && ib:
		v1, v2 := geom.VecBetweenPoints(vs[a.b], vs[a.a]), geom.VecBetweenPoints(vs[a.b], vs[a.c])
		vs[a.c].Y += yOfv2(v1, v2)
	case ib && ic:
		v1, v2 := geom.VecBetweenPoints(vs[a.b], vs[a.c]), geom.VecBetweenPoints(vs[a.b], vs[a.a])
		vs[a.a].Y += yOfv2(v1, v2)
	}
	shape.Vertices = vs
	return shape
}

func (a Angle) Draw(img *image.RGBA, shape geom.Shape) {
}

func (a Angle) IncrementOver(i int, roll goshape.Roll) goshape.Relation {
	return a
}

func (a Angle) GetMiddle(shape geom.Shape) geom.Point {
	vs := shape.Vertices
	v1, v2 := geom.VecBetweenPoints(vs[a.b], vs[a.c]), geom.VecBetweenPoints(vs[a.b], vs[a.a])
	res := v1.Add(v2).TimesScalar(0.3)
	return vs[a.b].MoveByVector(res)
}

func (a Angle) Collide(pts ...int) bool {
	return false
}

func (a Angle) GetPoints() []int {
	return []int{a.a, a.b, a.c}
}
