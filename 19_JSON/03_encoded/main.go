package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{"Kerub Querubin", 22}

	// NewEncoder() gives out a pointer to encoder
	// NewEncoder() needs a writer to give out a pointer to encoder
	// os.Stdout is a writer
	// when we have a pointer to encoder, we can then use Encode()
	json.NewEncoder(os.Stdout).Encode(p1)

}
