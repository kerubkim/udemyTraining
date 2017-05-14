package main

import "fmt"

func main() {

	check := true

	if food := "Chocolate"; check == true { // initialization ; expression
		// declared and assigned food; if check is equal to true
		fmt.Println(food) // print food
		// variable food is only accesible from this scope
	}

}
