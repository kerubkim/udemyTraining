package main

import "fmt"

func main() {
	x := average(1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(x)
}

// Variadic function
// adding "..." means you can add as many arguments in a function
func average(num ...float64) float64 {
	var total float64
	for _, n := range num { // using range to loop through num, range returns index/key and value
		total += n // replaced index/key with _ bcuz index/key is not needed here
	}
	return total / float64(len(num))
}
