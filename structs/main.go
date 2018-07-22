package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	alex := person{
		firstName: "alex",
		lastName:  "alegre",
		contact: contactInfo{
			email:   "alex@mail.com",
			zipCode: 85001,
		},
	}

	fmt.Printf("%+v", alex)
}
