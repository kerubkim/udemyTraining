package main

import "fmt"

func wrapper() func() int { // declared a function called wrapper that returns a function, that returns an int
	x := 0			// declairing variable x
	return func() int { // returns a function that returns an int
		x++
		return x
	}
}

func main() {
	increment := wrapper() // assigned wrapper function into a variable called increment
	fmt.Println(increment()) // initiated increment
	fmt.Println(increment())
	fmt.Println(increment())
}
