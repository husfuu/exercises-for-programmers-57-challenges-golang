package main

import "testing"

// test case
// area = 360
// area = 700
// area = 300
// area = 350

type getAreaTest struct {
	width  float64
	length float64
	exp    float64
}

type getPaint2CoverTest struct {
	area float64
	exp  int
}

func TestGetArea(t *testing.T) {
	var width1 float64 = 175
	var length1 float64 = 2
	var exp1 float64 = 350
	var tests = []getAreaTest{
		{
			width1,
			length1,
			exp1,
		},
	}
	for _, test := range tests {
		output := getArea(test.width, test.length)

		if output != test.exp {
			t.Errorf("Output %f not equal to expected %f", output, test.exp)
		}
	}
}

func TestGetPaint2Cover(t *testing.T) {
	var area1 float64 = 350
	var exp1 int = 1
	var area2 float64 = 360
	var exp2 int = 2
	var area3 float64 = 700
	var exp3 int = 2
	var area4 float64 = 0
	var exp4 int = 0
	var tests = []getPaint2CoverTest{
		{
			area1,
			exp1,
		},
		{
			area2,
			exp2,
		},
		{
			area3,
			exp3,
		},
		{
			area4,
			exp4,
		},
	}
	for _, test := range tests {
		output := getPaint2Cover(test.area)

		if output != test.exp {
			t.Errorf("Output %d not equal to expected %d", output, test.exp)
		}
	}
}
