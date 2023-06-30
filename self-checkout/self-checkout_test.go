package main

import "testing"

type getSubtotalTest struct {
	qty   int
	price float64
	exp   float64
}

type getTaxTest struct {
	subtotal float64
	exp      float64
}

type getTotalTest struct {
	subtotal float64
	tax      float64
	exp      float64
}

func TestGeSubtotalTest(t *testing.T) {
	// var qty1 = 6
	// var price1 = 5
}
