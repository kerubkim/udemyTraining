package main

import "fmt"

func main() {

	for i := 0; i < 100; i++ {
		fmt.Printf("%v \t %v \t %v \n", i, string(i), []byte(string(i)))
	}

}
