package main

import "fmt"

// type -> Declaration
// person -> Type
// struct -> underlying type
/// firstName -> Fields
/// string -> Known type

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	// Creating a variable to type person

	// := -> shorthand declaration
	// {"Kerub", "Querubin", 22} -> Struct literal

	p1 := person{"Kerub", "Querubin", 22}
	p2 := person{"Lexi", "Bae", 4}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println("First Name : ", p1.firstName, "Last Name : ", p1.lastName, "Age : ", p1.age)

}
