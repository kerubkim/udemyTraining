package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// Adds two thread
	wg.Add(2)
	go foo()
	go bar()
	// waits for two threads to finish
	wg.Wait()
}

func foo() {
	for i := 0; i < 15; i++ {
		fmt.Println("Foo:", i)
		// adding a pause
		time.Sleep(3 * time.Millisecond)
	}
	// tells wg this thread is finish
	wg.Done()
}

func bar() {
	for i := 0; i < 15; i++ {
		fmt.Println("Bar:", i)
		// adding a pause
		time.Sleep(3 * time.Millisecond)
	}
	// tells wg this thread is finish
	wg.Done()
}
