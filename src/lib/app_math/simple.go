package app_math

func NumberInRange[NT float32 | float64 | int](rMin, rMax, tMin, tMax, m NT) NT {
	return ((m-rMin)/(rMax-rMin))*(tMax-tMin) + tMin
}
