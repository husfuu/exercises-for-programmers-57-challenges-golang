package main

import "math"

// assume 1 galon -> 350 m^2

func getArea(width float64, length float64) float64 {
	return width * length
}

func getPaint2Cover(area float64) int {
	const areapergalon = 350
	area2cover := area / areapergalon
	return int(math.Ceil(area2cover))
}
