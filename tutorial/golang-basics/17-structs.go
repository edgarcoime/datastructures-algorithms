package main

import "fmt"

// ----- STRUCTS -----
// Structs allow you to store values with many
// data type

type Customer struct {
	name    string
	address string
	balance float64
}

func getCustInfo(c Customer) {
	fmt.Printf("%s owes us %.2f\n", c.name, c.balance)
}

func newCustAdd(c *Customer, address string) {
	c.address = address
}

func customerExample() {
	var cust Customer
	cust.name = "John Smith"
	cust.address = "5 main street"
	cust.balance = 523.45
	getCustInfo(cust)
	newCustAdd(&cust, "123 South St")
	cust2 := Customer{"Sally Smith", "123 Main", 0.0}
	println("Name:", cust2.name)
}

// Following functinos are part of Rectangle struct
type Rectangle struct {
	length, height float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.height
}

func rectangleExample() {
	rect1 := Rectangle{10.0, 15.0}
	println("Rect Area:", rect1.Area())
}

func main() {
	customerExample()
	rectangleExample()
}
