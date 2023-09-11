package app_math

func Clampf32(value, min, max float32) float32 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}

	return value
}
