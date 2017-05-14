package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a) // Prints the variable
	fmt.Print(&a)  // Prints the memory address of a

	var b *int = &a // a new variable that is a pointer to an int, and assigned &a
	// "*int" = pointing to a memory address where "int" is stored
	// b is type of INT pointer

	fmt.Println(b)
}
