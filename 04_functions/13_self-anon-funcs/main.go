package main

import "fmt"

func main() {
	func() { // annonymus because there is no name for this function
		fmt.Println("This is a self executing annomys function")
	}() // having "()" at the end of a functions makes it execute right away
}
