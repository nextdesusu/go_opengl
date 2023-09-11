package app_math

import "math"

type Vec2 struct {
	X, Y float32
}

func (v Vec2) ScalarMult(by float32) Vec2 {
	return Vec2{
		X: v.X * by,
		Y: v.Y * by,
	}
}

func (v Vec2) ScalarAdd(by float32) Vec2 {
	return Vec2{
		X: v.X + by,
		Y: v.Y + by,
	}
}

func (v Vec2) ScalarSub(by float32) Vec2 {
	return Vec2{
		X: v.X - by,
		Y: v.Y - by,
	}
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

func (v Vec2) Negate() Vec2 {
	return Vec2{
		X: -v.X,
		Y: -v.Y,
	}
}

func (v Vec2) Length() float32 {
	squaredX := v.X * v.X
	squaredY := v.Y * v.Y

	return float32(math.Sqrt(float64(squaredX + squaredY)))
}
