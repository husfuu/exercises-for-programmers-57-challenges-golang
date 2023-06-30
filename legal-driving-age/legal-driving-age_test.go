package main

import "testing"

type isLegalDrivingAgeTest struct {
	age int
	exp bool
}

func TestGetArea(t *testing.T) {
	age1 := 15
	age2 := 17
	age3 := 20
	exp1 := false
	exp2 := false
	exp3 := true
	var tests = []isLegalDrivingAgeTest{
		{
			age1,
			exp1,
		},
		{
			age2,
			exp2,
		},
		{
			age3,
			exp3,
		},
	}
	for _, test := range tests {
		output := isLegalDrivingAge(test.age)

		if output != test.exp {
			t.Errorf("Output %t not equal to expected %t", output, test.exp)
		}
	}
}
