package main

import "fmt"

type Contact struct {
	phone, address string
}

type Employee struct {
	name      string
	FirstName string
	LastName  string
	contact   Contact
}

func fullName(firstName string, lastName string) (fullName string) {
	fullName = firstName + " " + lastName
	return fullName
}

func (e Employee) enhancedFullName() string {
	return e.FirstName + " " + e.LastName
}

func (e *Employee) changeName(newName string) {
	(*e).name = newName
}

func (e *Employee) changePhone(newPhone string) {
	e.contact.phone = newPhone
}

func main() {
	e := Employee{
		FirstName: "Subham",
		LastName:  "Bhattacharjee",
	}

	fmt.Println(fullName(e.FirstName, e.LastName))
	fmt.Println(e.enhancedFullName())
	e.changeName("lol")
	e.changePhone("9046789016")
	fmt.Println(e.name)
	fmt.Println(e.contact.phone)
}
