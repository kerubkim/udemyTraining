package main

import "fmt"

func main() {
	// a slice that has length of 3, after 3 append is needed to add to mySlice
	mySlice := make([]int, 3, 5)                   // make a slice of - type int, starting with 0 elements and the length of (0) , with starting capacity of 5
	fmt.Println("Slice - ", mySlice)               // Prints mySlice
	fmt.Println("Slice Length - ", len(mySlice))   // Print the length of mySlice
	fmt.Println("SLice Capacity - ", cap(mySlice)) // Print the max capacity of mySlice

	fmt.Println("--------SECOND EXAMPLE-------")
	mySliceTwo := make([]int, 3) // if only 2 arguments is passed in to make() , the second argument will be the length and capacity of the slice
	mySliceTwo[0] = 123 // you can also add elements to an index this way
	mySliceTwo[1] = 1234
	mySliceTwo[2] = 12345
	//mySliceTwo[3] = 123456 // this would error out with "index out of range"
	//to add an element that exceeds the length of the slice
	mySliceTwo = append(mySliceTwo,123456) // adding the element pass the starting length
	fmt.Println("From mySliceTwo", mySliceTwo)
}