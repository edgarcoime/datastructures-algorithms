package main

import "fmt"

type Contact struct {
	fName string
	lName string
	phone string
}

type Business struct {
	name    string
	address string
	Contact
}

func (b Business) info() {
	fmt.Printf("Contact at %s is %s %s\n", b.name, b.Contact.fName, b.Contact.lName)
}

func main() {
	con1 := Contact{
		"James",
		"Wang",
		"555-1212",
	}
	bus1 := Business{
		"ABC Plumbing",
		"234 North St",
		con1,
	}
	bus1.info()
}
