package main

import "fmt"

func main() {
	i := 0
	for { // a for loop without a condition
		fmt.Println(i)
		if i >= 5 { // a condition if i is greater than 5, run this
			break // exit out of the for loop
		}
		i++ // increment i
	}
}
