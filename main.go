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

	anupam.printPerson()
}

func (p Person) printPerson() {
	fmt.Println(p)
}
