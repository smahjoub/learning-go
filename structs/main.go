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

	var p2 person

	p2.firstName = "Alex"
	p2.lastName = "Andersson"

	fmt.Printf("%+v", p2)
}
