package main

import "testing"

type getSliceTest struct {
	arg      int
	expected int
}

type getLeftOverSliceTest struct {
	slice    int
	people   int
	expected int
}

func TestSlicePizza(t *testing.T) {
	arg1 := 2     // pizza number
	expect1 := 16 // slice of all pizza

	var tests = []getSliceTest{
		{
			arg1,
			expect1,
		},
	}
	for _, test := range tests {
		output := getSlicePizza(test.arg)
		if output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}

func TestLeftOverSlice(t *testing.T) {
	slice1 := 16
	people1 := 8
	expect1 := 0

	var tests = []getLeftOverSliceTest{
		{
			slice1,
			people1,
			expect1,
		},
	}
	for _, test := range tests {
		output := getLeftOverSlice(test.slice, test.people)
		if output != test.expected {
			t.Errorf("Output %d not equal to expected %d", output, test.expected)
		}
	}
}
