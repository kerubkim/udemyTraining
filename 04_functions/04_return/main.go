package main

import "fmt"

func main() {
	fmt.Println(greet("Whats sup", "Kerub"))
}

func greet(greeting, name string) string {
	return fmt.Sprint(greeting, name) // Sprint - concats strings
}
