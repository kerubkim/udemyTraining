package main

import "fmt"

// closing a channel prevents adding to the channel
// closed channel can still recieve values from the channel

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			// this function is on hold until the value gets passed and printed out from the range loop
			c <- i
		}
		// closing unbuffered channel " c "
		close(c)
	}()

	// recieves values from the channel and assign it to n
	// loop through the channel until channel is closed and empty
	for n := range c { // range is really helpfull with channels
		fmt.Println(n)
	}
}
