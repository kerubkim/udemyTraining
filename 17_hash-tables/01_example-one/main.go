package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// make a http GET request to the url that gives a
	// response(res) and an error(err)
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	// if there is an error log errors
	if err != nil {
		log.Fatalln(err)
	}
	// ioutil.ReadAll reads what is passed in and return the data and an error if error exists
	// saves the data returned to variable bs
	// err is where the err data lives if there is an error
	// if there is no error the variable err would be nil
	bs, err := ioutil.ReadAll(res.Body)
	// if err is not nil log the error
	if err != nil {
		log.Fatalln(err)
	}
	// downcast bs to a string and save it to a variable str
	str := string(bs)
	// print str
	fmt.Println(str)

}
