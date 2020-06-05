package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateFirstName(firstName string) {
	(*p).firstName = firstName
}

func main() {
	// p := person{"Saif Eddine", "Mahjoub"}

	p := person{
		firstName: "Saif Eddine",
		lastName:  "Mahjoub",
		contactInfo: contactInfo{
			email:   "se.mahjoub@mail.com",
			zipCode: "75000",
		},
	}

	pp := &p
	pp.updateFirstName("Saif")

	p.print()

	var p2 person

	p2.firstName = "Alex"
	p2.lastName = "Andersson"
	p2.contactInfo = contactInfo{
		email:   "se.mahjoub@mail.com",
		zipCode: "75000",
	}

	p2.print()
}
