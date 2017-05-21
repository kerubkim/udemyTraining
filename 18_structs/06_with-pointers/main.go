package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// adding & - pointer
	p1 := &Person{"James Bond", 40}
	fmt.Println(p1)          // print p1
	fmt.Printf("%T \n", p1) // print type of p1
	fmt.Println(p1.Name)     // print p1 Name
}
