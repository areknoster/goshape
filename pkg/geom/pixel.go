package geom

type Pixel struct {
	X, Y int
}

func (pp Pixel) MoveByVec(v Pixel) Pixel {
	return Pixel{pp.X + v.X, pp.Y + v.Y}
}
