package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	var p1 Person
	// making a new reader
	rdr := strings.NewReader(`{"Name":"Kerub Querubin","Age": 22}`) // `{"Name":"Kerub Querubin","Age": 22}` an example of what you'd get from an API call

	// NewDncoder() gives out a pointer to decoder
	// NewDncoder() needs a reader to give out a pointer to decoder
	// when we have a pointer to decoder, we can then use Decode()

	// rdr (response from an API) is put inside of pointer &p1
	json.NewDecoder(rdr).Decode(&p1)

	fmt.Println(p1.Name)
	fmt.Println(p1.Age)
}
