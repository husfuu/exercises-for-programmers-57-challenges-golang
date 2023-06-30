package main

import (
	"fmt"
	"time"
)

func getAge() int {
	var age int
	fmt.Print("What is your current age? ")
	fmt.Scan(&age)
	return age
}

func getRetAge() int {
	var retAge int
	fmt.Print("At what age would you like to retire? ")
	fmt.Scan(&retAge)
	return retAge
}

func retCalc(age int, retAge int) string {
	diff := retAge - age
	curYear := time.Now().Year()
	retYear := curYear + diff
	res := fmt.Sprintf("You have %d years left until you can retire. It's %d, so you can retire in %d.", diff, curYear, retYear)
	return res
}

func main() {
	age := getAge()
	retAge := getRetAge()
	fmt.Println(retCalc(age, retAge))
}
