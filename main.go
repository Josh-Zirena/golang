package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := person{
		firstName: "Jim", lastName: "Jackson", contactInfo: contactInfo{
			email:   "JimJackson@gmail.com",
			zipCode: 33155,
		},
	}

	jim.updateName("Billy")
	jim.print()
}

// Golang always passes by value - copying the value
func (pointerToPerson *person) updateName(firstName string) {
	(*pointerToPerson).firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
