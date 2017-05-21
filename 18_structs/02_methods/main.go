package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

// Methods are a function inside of type struct
// (p person) -> Reciever
// fullName() -> Function name
// string -> Return type

// any type person has access to fullName(), because of the reciever (p person)
func (p person) fullName() string {
	return p.firstName + p.lastName
}

func main() {
	p1 := person{"Kerub", "Querubin", 22}
	fmt.Println("Hello my name is", p1.fullName())
}
