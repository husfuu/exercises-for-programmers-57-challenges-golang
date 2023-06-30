package main

import "fmt"

func main() {
	question := "What is the input string?"
	var name string
	for {
		fmt.Println(question)
		// taking user input
		fmt.Scanln(&name)
		if len(name) == 0 {
			fmt.Println("user must enter the name!")
			continue
		}
		result := fmt.Sprintf("%s has %d characters", name, len(name))
		fmt.Println(result)
		break		 		
	}
}