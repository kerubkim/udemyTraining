package main

import "fmt"

func main() {
	for i := 0; i < 11; i++ { // outerloop
		for j := 0; j < 6; j++ { // per 1 outerloop iteration this innerloop runs 5 times
			fmt.Println(i, " - ", j)
		}
	}
}
