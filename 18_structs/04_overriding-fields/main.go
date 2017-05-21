package main

import "fmt"


//Overriding does not replace the initial value, it is still accessible
// by direct notation


// Person Name -> Inner Type
type Person struct {
	Name string
	Age  int
}

// Hitman Name -> Outer Type
type Hitman struct {
	Person
	Name    string // This will override the inner Person Name
	License bool
}

func main() {
	p1 := Hitman{
		Person: Person{
			Name: "This will be overwritten",
			Age:  40,
		},
		Name:    "Overrided Inner Person Name",
		License: true,
	}
	fmt.Println("New name : ",p1.Name)
	// accessing by direct notation
	fmt.Println("Old name : ",p1.Person.Name)
}
