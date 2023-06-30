package main

import "testing"

type getSimpleInterestTest struct {
	pA  float64
	r   float64
	t   int
	exp float64
}

func TestGetArea(t *testing.T) {
	var pA1 float64 = 1500
	var r1 float64 = 4.3 / 100
	var t1 int = 4
	var exp1 float64 = 1758
	var tests = []getSimpleInterestTest{
		{
			pA1,
			r1,
			t1,
			exp1,
		},
	}
	for _, test := range tests {
		output := getSimpleInterest(test.pA, test.r, test.t)

		if output != test.exp {
			t.Errorf("Output %f not equal to expected %f", output, test.exp)
		}
	}
}
