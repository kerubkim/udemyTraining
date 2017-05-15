package main

import "fmt"

func main() {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 0} // declaring data with an array/slice of float64 numbers
	x := average(data...) // since data is not type float64 adding "..." makes it passable to the function and the float64's inside of data is accessible by the function
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
