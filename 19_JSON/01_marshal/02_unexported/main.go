package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	name string
	age  int
}

func main() {
	p1 := Person{"James Bond", 40}
	bs, _ := json.Marshal(p1)
	fmt.Println("p1 is : ", p1)
	fmt.Println("JSON : ", string(bs))
	fmt.Println("JSON is empty because fields in Person are not exported (lower cased)")
}
