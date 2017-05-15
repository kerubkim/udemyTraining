package main

import "fmt"

func main() {
	// assigning a function inside of a variable
	// a way of having a function inside of a function
	x := func() {
		fmt.Println("Hello World")
	}
	x()
}
