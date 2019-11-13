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
	alex := person{firstName: "Alex", lastName: "Jordan"}
	// alexPointer := &alex
	// (&alex).updateName("Michael")
	// alexPointer.updateName("Michael")
	alex.updateName("Michael")
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
