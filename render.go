package goshape

import (
	"github.com/sirupsen/logrus"

	"goshape/pkg/geom"
)

func Bresenham(a, b geom.PixelPoint, putPixel func(point geom.PixelPoint)) {
	if a.X > b.X {
		a, b = b, a
	}
	logrus.Debugf("drawing line: %v -> %v", a, b)
	dx := b.X - a.X
	dy := b.Y - a.Y
	if dy < 0 {
		dy = -dy
	}

	switch {
	case a == b: //point
		putPixel(a)
	case dx == 0: //vertical
		if a.Y > b.Y {
			a, b = b, a
		}
		for ; a.Y <= b.Y; a.Y++ {
			putPixel(a)
		}
	case dy == 0: //horizontal
		if a.X > b.X {
			a, b = b, a
		}
		for ; a.X <= b.X; a.X++{
			putPixel(a)
		}
	case dx > dy:
		if a.Y < b.Y {
			dy, e, slope := 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				putPixel(a)
				a.X++
				e -= dy
				if e < 0 {
					a.Y++
					e += slope
				}
			}
		} else {
			// BresenhamDxXRYU(img, x1, y1, x2, y2, col)
			dy, e, slope := 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				putPixel(a)
				a.X++
				e -= dy
				if e < 0 {
					a.Y--
					e += slope
				}
			}
		}
		putPixel(b)
	case dy >= dx:
		if a.Y < b.Y {
			dx, e, slope := 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				putPixel(a)
				a.Y++
				e -= dx
				if e < 0 {
					a.X++
					e += slope
				}
			}
		} else {
			// BresenhamDyXRYU(img, x1, y1, x2, y2, col)
			dx, e, slope := 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				putPixel(a)
				a.Y--
				e -= dx
				if e < 0 {
					a.X++
					e += slope
				}
			}
		}
		putPixel(b)

	}
}