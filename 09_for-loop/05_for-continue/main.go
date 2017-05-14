package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 { // if "i" does NOT have a remainder, continue to the next iteration
			continue
		}
		// after this point it skips, if it is even
		fmt.Println(i)
		if i >= 11 {
			break
		}
	}
}
