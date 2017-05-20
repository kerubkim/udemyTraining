package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

// Counting how many words would be in a letter A - Z  from a book
func main() {
	// http GET request of a book
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	// BEFORE the function where defer lives finishes running, defer res.Body.Close() will run which closes res.Body
	defer res.Body.Close() // have to defer res.Body.Close() everytime you do a http request

	// split what was scanned into words
	scanner.Split(bufio.ScanWords)

	// Create a slice that would hold counts
	buckets := make([]int, 200)

	// loop through the words
	for scanner.Scan() {
		// n is the number that was returned from hashbucket function
		n := HashBucket(scanner.Text())
		// adds one to a bucket
		buckets[n]++

	}
	// print how many each letter exists in the book
	// a - z
	// 65 through 123  would be USCII numbers
	fmt.Println(buckets[65:123])

}

// a function that takes in a word and returns the first letter as an Int
func HashBucket(word string) int {
	// grabs the first letter of the word that was passed in
	return int(word[0])
}
