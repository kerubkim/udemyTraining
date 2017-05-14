package main

import "fmt"

func main() {

	if true {
		fmt.Println("This ran") // this will run on load
	}
	if false {
		fmt.Println("This did not run")
	}
	if !true {
		fmt.Println("This did not run !") // this will run on load
	}
	if !false {
		fmt.Println("This ran with !") // this will run on load
	}
}
