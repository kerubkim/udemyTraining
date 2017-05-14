package main

import "fmt"

func main() {

	for i := 0; i < 100; i++ {
		fmt.Printf("%v \t %v \t %v \n", i, string(i), []byte(string(i)))
		// taking "i" into a string, making "i" string to a slice of byte
	}
	foo := 'k' // using '' turns the characters to a rune
	fmt.Println(foo)
	fmt.Printf("%T \n", foo) // "%T" shows the type of "foo"
	fmt.Println(`
	<div>
		<h3>Test</h3>
	</div>
	`)
	// using ` ` instead of " " for strings keeps the spaces and tabs
	// if " " was used everything side would have to be in one line
}
