package main

import "fmt"

func check(n *int) func() string { // a function that takes memory address pointer as a parameter and returns a function
	return func() string {
		x := *n % 2 // check if there is a remainder
		if x == 1 { // if remainder is 1
			return "ITS ODD"
		}
		return "ITS EVEN"
	}

}

func main() {
	//x := 12 % 2 // for remainders use "%" not "/"
	//fmt.Println(x)
	//if x == 1 {
	//	fmt.Println("ODD")
	//}else{
	//	fmt.Println("EVEN")
	//}
	var num int
	fmt.Print("Enter a number : ") // ask client for a number
	fmt.Scan(&num) // save the input number into the variable "num"
	result := check(&num) // place the check function into a variable "result"
	fmt.Println(result()) // invoke the "result function"
}
