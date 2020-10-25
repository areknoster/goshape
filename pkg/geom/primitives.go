package geom

type Point struct {
	X, Y float64
}

type Vertex Point
type Vector Point

func NormPoint(px, py, w, h int) Point {
	return Point{
		X: float64(px)/float64(w),
		Y: 1.0 - float64(py)/float64(h),
	}
}

func DenormPoint(p Point, w,h int) (PixelPoint){
	return PixelPoint{int(float64(w) * p.X), int(float64(w) *(1.0 - p.Y))}
}

func MoveByVector(v Vertex, vec Vector) Vertex {
	return Vertex{
		X: v.X + vec.X,
		Y: v.Y + vec.Y,
	}
}

type PixelPoint struct {
	X, Y int
}

