package app_math

type Range2d struct {
	Min, Max float64
}

func NumberInRange2D(srcRange Range2d, targetRange Range2d, m float64) float64 {
	s1 := m - srcRange.Min/srcRange.Max - srcRange.Min
	s2 := targetRange.Max - targetRange.Min
	return s1*s2 + targetRange.Min
}
