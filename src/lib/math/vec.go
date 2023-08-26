package math

type Vec2 struct {
	X, Y float64
}

type Vec3 struct {
	Vec2,
	Z float64
}

var SPACE_3D = Range2d{
	Min: -1,
	Max: 1,
}

func (self Vec2) V2NormalizeToScreenSize(size Size) Vec2 {
	x := NumberInRange2D(SPACE_3D, Range2d{Min: 0, Max: size.Width}, self.X)
	y := NumberInRange2D(SPACE_3D, Range2d{Min: 0, Max: size.Height}, self.Y)

	return Vec2{
		X: x,
		Y: y,
	}
}
