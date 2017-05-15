package main

import "fmt"

// filterthis takes an slice/array of int and a function that takes in an int that returns a bool
// filterthis returns an array/slice of int
func filterthis(nums []int, callback func(int) bool) []int {
	var x []int
	for _, n := range nums {
		if callback(n) {
			x = append(x, n) // append n to the array/slice x
		}
	}
	return x
}
func main() {
	runme := filterthis([]int{1, 2, 3, 4, 5, 6, 7}, func(z int) bool {
		return z > 1
	})
	fmt.Println(runme)
}
