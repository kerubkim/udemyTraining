package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)

	fmt.Println(*b) // dereferencing by adding * in the begining of the variable

}
