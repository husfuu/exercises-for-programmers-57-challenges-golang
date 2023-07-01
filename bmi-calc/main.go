package main

import "fmt"

func calcBMI(weight float32, height float32) (float32, string) {
	bmi := (weight / (height * height))
	label := ""
	if bmi > 18.5 && bmi < 25 {
		label = "ideal"
	} else if bmi > 25 {
		label = "overweight"
	} else {
		label = "underweight"
	}

	return bmi, label
}

func getWeightAndHeight() (float32, float32) {
	var weight float32
	var height float32
	fmt.Printf("input your weight and height: ")
	fmt.Scanf("%f %f", &weight, &height)
	fmt.Printf("weight %f \nheight %f", weight, height)
	fmt.Println()
	return weight, height
}

func concl(bmi float32, label string) string {
	return fmt.Sprintf("Your BMI is %f. \nYou are %s. You should see your doctor.", bmi, label)
}

func main() {
	weight, height := getWeightAndHeight()
	bmi, label := calcBMI(weight, height)
	fmt.Println(concl(bmi, label))
}
