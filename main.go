package main

import (
	"fmt"
)

type ContactInfo struct {
	email   string
	pinCode string
}

type Person struct {
	firstname string
	lastname  string
	contact   ContactInfo
}

func main() {
	anupam := Person{
		firstname: "Anupam",
		lastname:  "Datta",
		contact: ContactInfo{
			email:   "anupam@gmail.com",
			pinCode: "734005",
		},
	}

	// update first name
	// can call a function of type pointer to Person
	anupam.updateFirstName("bapi")

	anupam.printPerson()
}

func (p *Person) updateFirstName(newFirstName string) {
	// *p: value at p, which is 'anupam' object
	// get whatever inside the memory address pointed by p
	(*p).firstname = newFirstName
}

func (p Person) printPerson() {
	fmt.Println(p)
}
