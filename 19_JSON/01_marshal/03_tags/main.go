package main

import (
	"encoding/json"
	"fmt"
)

// `json:"-"` and `json:"wisdom score"`  -> are tags.
// `json:"-"` -> this field will be skipped where ever this tag is placed
// `json:"a new field name"` -> changes the name of the field when
type Person struct {
	Name       string
	Profession string `json:"-"`
	Age        int    `json:"wisdom score"`
}

func main() {
	p1 := Person{"James Bond", "Killer", 40}
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
	fmt.Println("Notice how the field Profession is skipped and Age field name turned to wisdome score, because of the use of tags")
}
