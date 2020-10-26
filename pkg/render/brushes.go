package render

import (
	"image"
	"image/color"

	"goshape/pkg/geom"
)

func RenderCommonShape(shape geom.Shape, img *image.RGBA, theme Theme) {
	w, h := img.Rect.Max.X, img.Rect.Max.Y
	put := NewCircleBrush(2, img, theme.LineColor)

	vs := shape.Vertices
	p1 := geom.DenormPoint(geom.Point(vs[len(vs)-1]), w, h)
	p2 := geom.DenormPoint(geom.Point(vs[0]), w, h)
	BresenhamLine(p1, p2, put) //between last and first
	for i := 1; i < len(vs); i++ {
		p1 = p2
		p2 = geom.DenormPoint(geom.Point(vs[i]), w, h)
		BresenhamLine(p1, p2, put)
	}
}

func RenderSelectedShape(shape geom.Shape, img *image.RGBA, theme Theme) {
	w, h := img.Rect.Max.X, img.Rect.Max.Y
	putVertex := NewFullCircleBrush(8, img, theme.LineColor)
	put := NewCircleBrush(2, img, theme.AccentColor)

	vs := shape.Vertices
	p1 := geom.DenormPoint(geom.Point(vs[len(vs)-1]), w, h)
	p2 := geom.DenormPoint(geom.Point(vs[0]), w, h)
	BresenhamLine(p1, p2, put) //between last and first
	for i := 1; i < len(vs); i++ {
		p1 = p2
		p2 = geom.DenormPoint(geom.Point(vs[i]), w, h)
		BresenhamLine(p1, p2, put)
	}

	for _, vertex := range shape.Vertices {
		putVertex(geom.DenormPoint(geom.Point(vertex), w, h))
	}
}

func NewFullCircleBrush(r int, img *image.RGBA, color color.RGBA) Brush {
	if r == 0 {
		return func(pp geom.Pixel) {
			img.SetRGBA(pp.X, pp.Y, color)
		}
	}

	return func(pp geom.Pixel) {
		d := 2*r - 1
		rsq := r * r
		shift := r - 1
		start := geom.Pixel{pp.X - shift, pp.Y - shift}
		for i := 0; i < d; i++ {

			for j := 0; j < d; j++ {
				cur := geom.Pixel{start.X + i, start.Y + j}
				dx, dy := cur.X-pp.X, cur.Y-pp.Y
				if dx*dx+dy*dy < rsq {
					img.SetRGBA(cur.X, cur.Y, color)
				}
				cur.Y++
			}
		}
	}
}

func NewCircleBrush(r int, img *image.RGBA, color color.RGBA) Brush {
	if r == 0 {
		return func(pp geom.Pixel) {
			img.SetRGBA(pp.X, pp.Y, color)
		}
	}
	return func(pp geom.Pixel) {
		p := geom.Pixel{0, r}
		xsq := 0
		rsq := r * r
		ysq := rsq
		put := func(p geom.Pixel) {
			dest := geom.MoveByVec(pp, p)
			img.SetRGBA(dest.X, dest.Y, color)
		}
		for p.X <= p.Y {
			put(geom.Pixel{p.X, p.Y})
			put(geom.Pixel{p.Y, p.X})
			put(geom.Pixel{-p.X, p.Y})
			put(geom.Pixel{-p.Y, p.X})
			put(geom.Pixel{p.X, -p.Y})
			put(geom.Pixel{p.Y, -p.X})
			put(geom.Pixel{-p.X, -p.Y})
			put(geom.Pixel{-p.Y, -p.X})

			xsq = xsq + 2*p.X + 1
			p.X++
			// Potential new y^2 = (y-1)^2 = y^2 - 2y + 1
			y1sq := ysq - 2*p.Y + 1
			// Choose y or y-1, whichever gives smallest error
			a := xsq + ysq
			b := xsq + y1sq
			if a-rsq >= rsq-b {
				p.Y--
				ysq = y1sq
			}
		}
	}

}

type Brush func(pp geom.Pixel)

func NewSquareBrush(radius int, img *image.RGBA, color color.RGBA) Brush {
	if radius == 1 {
		return func(pp geom.Pixel) {
			img.SetRGBA(pp.X, pp.Y, color)
		}
	}
	r := radius - 1
	return func(pp geom.Pixel) {
		vs := []geom.Pixel{
			{pp.X - r, pp.Y - r},
			{pp.X - r, pp.Y + r},
			{pp.X + r, pp.Y + r},
			{pp.X + r, pp.Y - r},
		}

		dim := 2 * r
		for _, v := range vs { //corners
			img.SetRGBA(v.X, v.Y, color)
		}
		for i := 1; i < dim; i++ {
			img.SetRGBA(vs[0].X+i, vs[0].Y, color)
			img.SetRGBA(vs[1].X+i, vs[1].Y, color)
			img.SetRGBA(vs[2].X, vs[2].Y+i, color)
			img.SetRGBA(vs[3].X, vs[3].Y+i, color)
		}
	}
}


