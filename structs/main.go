package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{
		firstName: "alex",
		lastName:  "alegre",
		contactInfo: contactInfo{
			email:   "alex@mail.com",
			zipCode: 85001,
		},
	}

	alex.print()
	alex.updateName("alexander")
	alex.print()
}

func (p person) updateName(firstName string) {
	p.firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
	// print empty line
	fmt.Println()
}
