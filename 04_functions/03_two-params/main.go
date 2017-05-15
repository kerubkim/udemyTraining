package main

import "fmt"

func main() {
	greet("Yooooo", "kerub")
	greet2("Suppp", "K")
}

func greet(greeting string, name string) {
	fmt.Println(greeting, name)
}
// or if a function has params of the same type
func greet2(greeting, name string) {
	fmt.Println(greeting, name)
}
