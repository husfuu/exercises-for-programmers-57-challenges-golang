package main

import (
	"fmt"
)

func getLength() float64 {
	var length float64
	fmt.Print("What is the length of the room in feet? ")
	fmt.Scan(&length)
	return length
}

func getWidth() float64 {
	var width float64
	fmt.Print("What is the width of the room in feet? ")
	fmt.Scan(&width)
	return width
}

func areaOfRectangularRoom(length float64, width float64) string {
	area_sqr := length * width
	area_sqr_meter := area_sqr * 0.09290304
	result := fmt.Sprintf("The area is \n %.3f square feet \n %.3f square meters", area_sqr, area_sqr_meter)
	return result
}

func main() {
	length := getLength()
	width := getWidth()
	fmt.Println(areaOfRectangularRoom(length, width))
}
