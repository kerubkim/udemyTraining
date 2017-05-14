package main

import "fmt"

func SwitchOnType(x interface{}) { // interface{} means - any type, bcuz we dont know what is being passed in

	switch x.(type) { // this is an assert; asserting, "x is of this type"
	case int:
		fmt.Println("this is an int")
	case string:
		fmt.Println("this is a string")
	default:
		fmt.Println("Unknown type")
	}
}
func main() {
	SwitchOnType(42)
	SwitchOnType("kerub")
}
