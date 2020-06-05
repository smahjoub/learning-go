package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName   string
	lastName    string
	contactInfo contactInfo
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
	fmt.Println(p)

	var p2 person

	p2.firstName = "Alex"
	p2.lastName = "Andersson"
	p2.contactInfo = contactInfo{
		email:   "se.mahjoub@mail.com",
		zipCode: "75000",
	}
	fmt.Printf("%+v", p2)
}
