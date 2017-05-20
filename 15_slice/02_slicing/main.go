package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6}
	// print numbers starting at index 2 and index before 4
	fmt.Println(nums[2:4]) // prints [3 4]
	fmt.Println(nums[:5])  // prints from the begining of the slice to the 5th not including the 5th
	fmt.Println(nums[1:])  // prints from the 1th through the end of the slice
	fmt.Println(nums[:]) // prints everything inside of the slice
}
