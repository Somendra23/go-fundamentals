package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct {
}

type spanishBot struct {
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	fmt.Println("Interfaces in Action")
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	//custom logic generating eb bot
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}
