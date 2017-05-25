package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup

var counter int64

func main() {
	wg.Add(2)
	go increment("Foo")
	go increment("Bar")
	wg.Wait()
	fmt.Println("Final Counter", counter)
}

func increment(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(2)) * time.Millisecond)
		// atomic -> similiar to mutex, (pointer,number)
		atomic.AddInt64(&counter, 1) //increment pointer value by 1
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}

// To check for Race Conditions
// go run -race main.go
