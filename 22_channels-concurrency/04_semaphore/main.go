package main

import (
	"fmt"
)

// Another way of WaitGroup channel with out WaitGroup

func main() {
	c := make(chan int)
	done := make(chan bool)

	// ========== SELF EXECUTING FUNCTIONS ===========
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done   // waiting for a bool to be passed to the channel, then remove from the channel
		<-done   // waiting for a bool to be passed to the channel, then remove from the channel
		close(c) // once bools are passed , close the channel
	}()
	// ================================================
	for n := range c {
		fmt.Println(n)
	}
}
