package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	//imaginary custom logic for which generates greeting
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	//imaginary custom logic which generates greeting
	return "Hola!"
}
