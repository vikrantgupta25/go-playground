package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// very similar to type cards []string
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) updateName() {
	p.firstName = "Updated Name!"
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func main() {
	// can omit the properites as well but not recommended
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// var alex person
	// fmt.Println(alex) // zero value

	// // better way of logging in a struct with property names
	// fmt.Printf("%+v", alex)

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	alex := person{firstName: "alex", lastName: "anderson", contactInfo: contactInfo{
		email:   "alex@gmail.com",
		zipCode: 94000,
	}}

	alex.updateName()
	alex.print()

}
