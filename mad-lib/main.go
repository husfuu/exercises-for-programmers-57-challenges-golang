package main

import "fmt"

func getNoun(word string) string {
	fmt.Print("Enter a noun: ")
	fmt.Scan(&word)
	return word
}

func getVerb(word string) string {
	fmt.Print("Enter a verb: ")
	fmt.Scan(&word)
	return word
}

func getAdjective(word string) string {
	fmt.Print("Enter an adjective: ")
	fmt.Scan(&word)
	return word
}

func getAdverb(word string) string {
	fmt.Print("Enter an adverb: ")
	fmt.Scan(&word)
	return word
}

func getStory(noun string, verb string, adj string, adv string) string {
	return fmt.Sprintf("Do you %s your %s %s %s? That's hilarious!", verb, adj, noun, adv)
}

func main() {
	var word string
	noun := getNoun(word)
	verb := getVerb(word)
	adj := getAdjective(word)
	adv := getAdverb(word)

	fmt.Println(getStory(noun, verb, adj, adv))
	// example output
	// Do you walk your blue dog quickly? That's hilarious!
}
