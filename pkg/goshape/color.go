package goshape

import (
	"image/color"
)

func ColorToRGBA(col color.Color) color.RGBA{
	r, g, b, a := col.RGBA()
	return color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(a),
	}
}
