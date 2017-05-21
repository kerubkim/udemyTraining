package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// When using a Hitman type, Person fields needs to be defined.

type Hitman struct {
	Person  // Type Person promoted to type Hitman
	License bool
}

func main() {
	p1 := Hitman{
		Person: Person{
			Name: "James Bond",
			Age:  40,
		},
		License: true,
	}

	fmt.Println(p1.Name, p1.Age, p1.License)
}
