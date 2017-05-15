package main

import "fmt"

func main() {

	showme([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(x int) { // passing a slice/array of int numbers and an annonyms function
		fmt.Println(x)
	})

}

func showme(numbers []int, secondfunc func(int)) {
	for _, n := range numbers {
		secondfunc(n)
	}
}
// callbacks = a function that takes a function as a parameter