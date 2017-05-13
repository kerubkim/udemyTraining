package main

import "fmt"

func main() {
	x := 0                    // declaring and assigning x
	increment := func() int { // declaring an annonyms function inside variable called increment that returns an int
		x++
		return x
	}
	fmt.Println(increment()) // using the function increment inside of fmt
	fmt.Println(increment())
	fmt.Println(increment())
}
