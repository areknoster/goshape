package geom

func NormPoint(point Pixel, w, h int) Point {
	return Point{
		X: float64(point.X)/float64(w),
		Y: 1.0 - float64(point.Y)/float64(h),
	}
}

func NormVector(vec Pixel, w,h int ) Vector {
	return Vector{
		X: float64(vec.X)/float64(w),
		Y: -float64(vec.Y)/float64(h),
	}
}

func DenormPoint(p Point, w,h int) (Pixel){
	return Pixel{int(float64(w) * p.X), int(float64(w) *(1.0 - p.Y))}
}



