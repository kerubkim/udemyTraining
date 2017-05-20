package main

import "fmt"

func main() {
	// literal declaration of a map
	myMap := map[string]string{
		"greeting":  "hello",
		"greeting2": "good afternoon",
	}
	myMap["greeting3"] = "good bye"

	fmt.Println(myMap)
	fmt.Println(myMap["greeting3"])
}
