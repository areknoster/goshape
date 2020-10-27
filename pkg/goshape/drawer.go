package goshape

import (
	"image"

	"goshape/pkg/geom"
)

type Drawer interface{
	Draw(img *image.RGBA, shape geom.Shape)
}
