package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)
// instead of err, err is replaced with _
// when using _ if there is an error, it would throw away(skip) the error
func main() {
	res, _ := http.Get("https://www.google.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
