package main

import "fmt"

func main() {

	greetings := make(map[string]string)

	greetings["KeyOne"] = "Hello"
	greetings["KeyTwo"] = "Sup"
	greetings["KeyThree"] = "Whats Good"
	greetings["KeyFour"] = "High"

	// A range loop that loops through greetings
	// Range loop will give the key and value
	// the variables key and val can be named anything
	for key, val := range greetings {
		fmt.Println("Key is - ", key, " , Val is -", val)
	}
}
