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

	alex.updateName("alexander")
	alex.print()
}

func (pointerToPerson *person) updateName(firstName string) {
	(*pointerToPerson).firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
	// print empty line
	fmt.Println()
}
