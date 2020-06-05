package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type frenchBot struct{}

func main() {

	eb := englishBot{}
	fb := frenchBot{}

	fmt.Println(eb.getGreeting())
	fmt.Println(fb.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello!"
}

func (frenchBot) getGreeting() string {
	return "Bonjour!"
}
