package main

import "fmt"

func zero(z *int){ // a function that has a parameter of pointer to int
	*z = 0	// dereferencing and changing that memory address value to 0
}

func main(){

	x := 300 // making a variable
	zero(&x)  // passing the memory address of x as an argument to the function "zero"
	fmt.Println(x)

}
