package main

import "fmt"

func main() {

	myName := "Kerub"

	switch {
	case len(myName) == 2: // len(...) checks the length of myName which is 5
		fmt.Println("your name has 2 letters")
	case myName == "Querubin":
		fmt.Println("Yoo querubin")
	case myName == "Kim" , myName == "K":
		fmt.Println("Kim Kim Kim")
	default:
		fmt.Println("No cases matched")
	}

}
