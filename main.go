package main

import "fmt"

type Address struct {
	city string
	pin int
}
type Person struct {
	name string
	age int
	address Address
}
func main() {
	p1 := Person{
		name: "Lisa",
		age: 21,
		address: Address{
			city: "Kolkata",
			pin: 700078,
		},

	}
	fmt.Println(p1)
}
