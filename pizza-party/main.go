package main

import "fmt"

func getPeople() int {
	var people int
	fmt.Print("How many people? ")
	fmt.Scan(&people)
	return people
}

func getPizza() int {
	var pizza int
	fmt.Print("How many pizzas do you have? ")
	fmt.Scan(&pizza)
	return pizza
}

func getSlicePizza(pizza int) int {
	// get all slice of pizza that user have
	// assume one pizza have 8 slice
	const SLICE = 8
	return SLICE * pizza
}

func getLeftOverSlice(slice int, people int) int {
	return slice % people
}

func printPizzaParty(people int, pizza int, slice int, left_over int) string {
	res := fmt.Sprintf("%d people with %d pizzas \nEach person gets %d pieces of pizza. \nThere are %d leftover pieces.", people, pizza, slice, left_over)
	return res
}

func main() {
	people := getPeople()
	pizza := getPizza()
	slice := getSlicePizza(pizza)
	leftOver := getLeftOverSlice(slice, people)

	fmt.Println(printPizzaParty(people, pizza, slice, leftOver))
}
