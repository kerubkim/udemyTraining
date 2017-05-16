package main

import "fmt"

func hello() {
	fmt.Println("Hello ")
}
func world() {
	fmt.Println("World ")
}

func main() {
	// Defer - puts hold the function to be executed last
	defer world() // This function would run before the outer function exits
	hello()
}
