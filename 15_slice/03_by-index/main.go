package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(nums[2])    // print the value in index 2 (prints - 3)
	fmt.Println("Kerub"[3]) // print the ASCII number of the letter in index 3 of the word "Kerub" which is the ltter "u"
}
// strings are made of runs, runes are made of bytes, so string bytes - a slice of bytes