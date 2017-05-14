package main

import "fmt"

func main() {

	switch "Kerub" {
	case "K":
		fmt.Println("Its k") // this wont run
	case "Kerub":
		fmt.Println("ok kerub") // this will run
		fallthrough
	case "Kim":
		fmt.Println("lol kim") // this will run
		fallthrough
	case "Querubin":
		fmt.Println("querubin!") // this will run
	case "Something":
		fmt.Println("Something else") // this wont run
	default:
		fmt.Println("the D") // this wont run
	}

	// fallthrough - makes the next case to true / run
}
