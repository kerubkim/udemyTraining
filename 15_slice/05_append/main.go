package main

import "fmt"

func main() {

	mySlice := make([]int, 0, 5) // make a slice of ints, starting with no elements, and starting capacity of 5

	for i := 0; i < 15; i++ {
		mySlice = append(mySlice, i) // Appending to mySlice the value i
		fmt.Println("Slice - ", mySlice ,"Length - ", len(mySlice), mySlice, "Capacity - ", cap(mySlice), "Value at index ", i, " - ", i)
	}

}
