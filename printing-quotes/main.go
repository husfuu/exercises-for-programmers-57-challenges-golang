package main

import "fmt"

func getQuote(quote string) string {
	fmt.Println("What is the quote?")
	fmt.Scanln(&quote)
	return quote
}

func getAuthor(name string) string {
	fmt.Println("Who said it?")
	fmt.Scanln(&name)
	return name
}

func sayIt(name string, quotes string) string {
	return fmt.Sprintf("%s says, %s.", name, quotes)
}

func main() {
	var quotes string
	var name string

	quotes = getQuote(quotes)
	name = getAuthor(name)

	fmt.Println(sayIt(name, quotes))
}
