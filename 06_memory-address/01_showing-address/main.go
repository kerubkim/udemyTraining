package main

import "fmt"

func main() {
	x := 42
	fmt.Println("x - ", x)
	fmt.Println("x's memory address - ", &x) // the "&" allows us to look for x's memory address
}
