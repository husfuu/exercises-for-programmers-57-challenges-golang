package main

import "testing"

type calcTaxTest struct {
	state       string
	orderAmount float64
	exp         float64
}

func TestGetArea(t *testing.T) {
	var state1 string = "WI"
	var state2 string = "MN"
	var oA1 float64 = 10
	var exp1 float64 = 10.55
	var exp2 float64 = 10
	var tests = []calcTaxTest{
		{
			state1,
			oA1,
			exp1,
		},
		{
			state2,
			oA1,
			exp2,
		},
	}
	for _, test := range tests {
		output := calcTax(test.state, test.orderAmount)

		if output != test.exp {
			t.Errorf("Output %f not equal to expected %f", output, test.exp)
		}
	}
}
