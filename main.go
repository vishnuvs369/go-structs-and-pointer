package main

import "fmt"

type contactInfo struct {
	email    string
	postcode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	vishnu := person{
		firstName: "vishnu",
		lastName:  "vs",
		contactInfo: contactInfo{
			email:    "vichu@gmil.com",
			postcode: 38383,
		},
	}

	vishnu.updateName("unni")
	vishnu.Print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) Print() {
	fmt.Printf("%+v", p)
}
