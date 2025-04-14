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
	// contactInfo I could just use this, equivalent to "contactInfo contactInfo"
}

func main() {
	// alex := person{"Alex", "Anderson"} valid way but should not be used

	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Jhonson"
	// alex.contact.email = "alex@gmail.com"
	// alex.contact.zipCode = 3000
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 4000,
		},
	}

	// a way of using pointer. & give the memory address instead of the value
	// jimPointer := &jim
	// jimPointer.updateFirstName("Jimmy")
	// jimPointer.print()

	jim.print()
	jim.updateFirstName("Jimmy")
	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println("")
}

// using a pointer * to avoid passing arg as value. Gives the memory address' value
func (p *person) updateFirstName(name string) {
	// (*p).firstName = name This is the same as below, but () turns a memory pointer into its value

	p.firstName = name
}
