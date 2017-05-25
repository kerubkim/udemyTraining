package main

import (
	"fmt"
	"time"
)

// go func(){.....}() means creating a Go Routine
// channel <- value. Adding a value to the channel
// <- channel. Removes value from the channel

func main() {
	// making an unbuffered channel that will be able to communicate with int
	c := make(chan int)

	// creating a goroutine
	// waiting to pass
	go func() { // ============= func A
		for i := 0; i < 10; i++ {
			// placing i to the channel then stops until func B runs
			c <- i
		}
	}()

	// creating a goroutine
	// waiting to recieve
	go func() { // ============== func B
		for {
			// waiting for func A to add a number to the channel
			// taking the number if off the unbuffered channel
			fmt.Println(<-c) // from the channel take what ever value is there, also removes value from the channel
		}
	}()

	// time sleep waits one second before existing just to see the results for testing purpose
	time.Sleep(time.Second)
}
