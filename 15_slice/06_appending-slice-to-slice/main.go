package main

import "fmt"

func main() {

	sliceOne := []int{1, 2, 3, 4, 5, 6}
	sliceTwo := []int{7, 8, 9}

	// Appending sliceTwo to sliceOne
	sliceOne = append(sliceOne, sliceTwo...) // appending all the elements inside of sliceTwo
	sliceOne = append(sliceOne,sliceTwo[1]) // appending one element from sliceTwo
	fmt.Println("Two slice combined", sliceOne)

}
