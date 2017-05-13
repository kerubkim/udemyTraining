package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://www.google.com/") // makes a request to the url, returns a response and an error(if any)
	if err != nil {                                 // check if there is an error
		log.Fatal(err) // log the error
	}
	page, err := ioutil.ReadAll(res.Body) // reads the response and place in the variable page
	res.Body.Close() // must close the response body when finished
	if err != nil { // check if there is an error
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
}
