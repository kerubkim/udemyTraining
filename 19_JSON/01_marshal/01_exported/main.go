package main

import (
	"encoding/json"
	"fmt"
)

// inside of a type, capitalize fields are exported, NON capitalized fields are not exported
type Person struct {
	Name  string
	Age   int
	kills int // this wont be exported
}

func main() {
	p1 := Person{"Person One", 22, 1000}
	// bs -> byte slice
	// json.Marshal(p1) taking a struct and truning it into byte slice
	bs, _ := json.Marshal(p1)
	// print byte slice
	fmt.Println("Byte Slice : ", bs)
	// print type of byte slice
	fmt.Printf("Type : %T \n", bs)
	// print converted byte slice into a string
	fmt.Println("Byte Slice into String : ", string(bs))
	fmt.Println("Person kills is not included because its not exported")
}
