package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	englishBot := englishBot{}

	printGreeting(englishBot)

	spanishBot := spanishBot{}

	printGreeting(spanishBot)
}
