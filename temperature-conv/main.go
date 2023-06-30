package main

import "fmt"

func convFahrenheit2Celc(f float32) float32 {
	return (f - 32) * 5 / 9
}

func convCelc2Fahrenheit(c float32) float32 {
	return (c * 9 / 5) + 32
}

func getCelcOrFah() (float32, string) {
	var celc float32
	var frt float32
	var choice string
	fmt.Println("Press C to convert from Fahrenheit to Celsius. \nPress F to convert from Celsius to Fahrenheit.")
	fmt.Scan(&choice)
	fmt.Println("Your choice: ", choice)
	if choice == "C" {
		fmt.Print("Please enter the temperature in Fahrenheit: ")
		fmt.Scan(&celc)
		return celc, choice
	} else if choice == "F" {
		fmt.Print("Please enter the temperature in Celsius: ")
		fmt.Scan(&frt)
		return frt, choice
	}
	return 0, choice
}

func printRes(choice string, value float32) string {
	if choice == "C" {
		value = convFahrenheit2Celc(value)
	} else if choice == "F" {
		value = convCelc2Fahrenheit(value)
	}
	res := fmt.Sprintf("The temperature in %s is %f.", choice, value)
	return res
}

func main() {
	value, choice := getCelcOrFah()

	fmt.Println(printRes(choice, value))

}
