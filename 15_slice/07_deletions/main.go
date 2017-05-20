package main

import "fmt"

func main() {

	sliceOne := []string{"Monday", "Tuesday"}
	sliceTwo := []string{"Wednesday","Thursday", "Friday"}

	sliceOne = append(sliceOne,sliceTwo...)
	fmt.Println(sliceOne)
	sliceOne = append(sliceOne[:2], sliceTwo[1:]...) // it didnt add Wednesday (leaving out "Wednesday")

	fmt.Println(sliceOne)

}
