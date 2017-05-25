package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
func main() {
	// Adds how many thread
	wg.Add(2)
	go foo()
	go bar()
	// waits for threads to finish
	wg.Wait()
}

func foo() {
	for i := 0; i < 15; i++ {
		fmt.Println("Foo:",i)
	}
	// tells wg this thread is finish
	wg.Done()
}

func bar() {
	for i := 0; i < 15; i++ {
		fmt.Println("Bar:",i)
	}
	// tells wg this thread is finish
	wg.Done()
}
