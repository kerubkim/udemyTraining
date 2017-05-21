package main

import "fmt"

// * Outer type overrides Inner type

//Overriding does not replace the initial method, it is still accessible
// by direct notation

type Person struct {
	Name string
	Age  int
}

type Hitman struct { // <- Outer type
	Person  // <- Inner type
	License bool
}

func (p Person) Greeting() {
	fmt.Println("I am a method from Person")
}

func (h Hitman) Greeting() {
	fmt.Println("I am a method from Hitman")
}

func main() {
	p1 := Person{
		Name: "Regular Person",
		Age:  40,
	}
	p2 := Hitman{
		Person: Person{
			Name: "A Hitman",
			Age:  32,
		},
		License: true,
	}
	p1.Greeting() // This would use Person method because p1 is a Person type
	p2.Greeting() // This would use Hitman method because p2 is a Hitman type

	// using direct notation
	p2.Person.Greeting() // This would use Person method because p2 is still a person that is promoted to Hitman

}
