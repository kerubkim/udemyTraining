package main

import (
	"fmt"
	"encoding/json"
)

// can also use tags but not required
type Person struct {
	Name string
	Age  int `json:"wisdom score"` // if anythings gets passed in with the field name wisdom score, it would be saved in Age
}

func main() {
	// initializing a variable of type Person
	var p1 Person
	// all p1 fields would be empty
	fmt.Println(p1.Name)
	fmt.Println(p1.Age)

	fmt.Println("===================")

	// creating the data for p1 into a slice of bytes
	bs := []byte(`{"Name":"Kerub Querubin","wisdom score":22}`)
	// Umarshall analyzes(parses) bs and saves it to pointer of p1
	json.Unmarshal(bs,&p1)

	fmt.Println(p1.Name)
	fmt.Println(p1.Age)
	fmt.Printf("%T \n",p1)

}
