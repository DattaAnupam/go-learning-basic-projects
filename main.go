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

	// storing the memory address of 'anupam',
	// point to the memory address of 'anupam'
	// and is type of Person, as pointing to a data of type Person
	pAnupam := &anupam
	// update first name
	pAnupam.updateFirstName("bapi")

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
