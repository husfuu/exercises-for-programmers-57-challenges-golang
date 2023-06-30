package main

func getSimpleInterest(pA float64, r float64, t int) float64 {
	// pA: principal amount
	// r: annual rate
	// t: number of years
	return pA * (1 + r*float64(t))
}
