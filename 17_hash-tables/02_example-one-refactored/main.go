package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	words := make(map[string]string)

	// make a buffer by using bufio.
	// bufio is a buffer.
	// buffer is a location that holds data
	// buffer makes is it so the cpu is not tied up
	// buffer makes giving data faster by handing out items as soon as the computer needs it
	// bufio.NewScanner , scans res.body and save it to variable sc
	// sc is a buffer scanner because bufio was used
	sc := bufio.NewScanner(res.Body)

	// split the sc data by words
	sc.Split(bufio.ScanWords)

	// scans the entire list of words that is inside of sc
	for sc.Scan() {
		// each interation , gives out the actual word and place it inside of words map
		words[sc.Text()] = "" // its equal to nothing because we dont want the definition. just the words
	}

	// checking if there is an error that was returned by the scanner(variable sc)
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading input error: ", err)
	}

	// iterate through the map words
	// _ is for the definition of the words but we dont need definition on this example
	for w, _ := range words {
		fmt.Println(w)
	}

}
