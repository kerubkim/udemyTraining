package main

import "fmt"

func main() {

	var arr [12]byte // array that holds 12 values of bytes (length = 12)
	for i := 0; i < len(arr); i++ {
		arr[i] = byte(i) // convert i into bytes
	}
	fmt.Println(arr)

	for x, v := range arr {
		// prints the index, the type, value, value in binary, value in byte
		fmt.Printf("index - %v , type - %T , value - %v , binary - %b \n", x, v, v, v)
	}
}
