package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// p := person{"Saif Eddine", "Mahjoub"}

	p := person{firstName: "Saif Eddine", lastName: "Mahjoub"}
	fmt.Println(p)
}
