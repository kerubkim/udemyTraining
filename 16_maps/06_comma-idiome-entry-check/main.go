package main

import "fmt"

func main() {

	myMap := make(map[string]string)
	myMap["kOne"] = "Yooo I am key one"
	myMap["kTwo"] = "Sup I am key two and I am alive"
	myMap["kThree"] = "I am key three"
	myMap["kFour"] = "I am key three"

	// Toggle delete line to check if and else results
	//delete(myMap,"kTwo")

	// val is the placeholder for the key that is being passed in through the if statement for a check for the value
	// exists is a boolean that would return to true or false, depends on what the if state evaluates to
	// val and exists are only accessible in the if statement
	if val, exists := myMap["kTwo"]; exists {
		fmt.Println("The value does exists")
		fmt.Println("The value is - ", val)
		fmt.Println("Does the value exists? ", exists)
	} else {
		fmt.Println("The value does not exists")
		fmt.Println("The value is - ", val)
		fmt.Println("Does the value exists? ", exists)
	}

	fmt.Println("===== LITTLE EXTRA EXPERIMENT =====")
	// An if statement in one line
	fmt.Println("Does the two key value match? " , myMap["kFour"] == myMap["kThree"]) // returns a bool
}
