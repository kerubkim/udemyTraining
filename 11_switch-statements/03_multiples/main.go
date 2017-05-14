package main

import "fmt"

func main() {

	switch "Kerub" {
	case "K":
		fmt.Println("This is K")
	case "Kim", "Kerub":
		fmt.Println("This might be kim or kerub") // this will run because case its kim or kerub
	case "Querubin","Mee":
		fmt.Println("a Querubin or Mee")
	default:
		fmt.Println("Who are you?")
	}

}
