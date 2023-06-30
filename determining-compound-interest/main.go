package main

import "math"

func calcCompoundInterest(pA float64, r float64, n int, t int) float64 {
	// return pA * math.Pow(1+r/float64(n), float64(n*t))
	return math.Ceil(pA * math.Pow(1+r/float64(n), float64(n*t)))
}
