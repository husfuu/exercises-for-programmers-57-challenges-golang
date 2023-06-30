package main

import "testing"

type calcCompoundInterestTest struct {
	pA  float64
	r   float64
	n   int
	t   int
	exp float64
}

func TestCalcCompoundInterest(t *testing.T) {
	var pA1 float64 = 1500
	var r1 float64 = 4.3 / 100
	var t1 int = 6
	var n1 int = 4
	var exp1 float64 = 1939
	var tests = []calcCompoundInterestTest{
		{
			pA1,
			r1,
			n1,
			t1,
			exp1,
		},
	}
	for _, test := range tests {
		output := calcCompoundInterest(test.pA, test.r, test.n, test.t)

		if output != test.exp {
			t.Errorf("Output %f not equal to expected %f", output, test.exp)
		}
	}
}
