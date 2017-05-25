package main

import (
	"sync"
	"time"
	"math/rand"
	"fmt"
)

var wg sync.WaitGroup

// mutex -> a tool to prevent Race Conditions(over lapping processes)
var mutex sync.Mutex

var counter int

func main(){
	wg.Add(2)
	go increment("Foo")
	go increment("Bar")
	wg.Wait()
	fmt.Println("Final Counter",counter)
}

func increment(s string){
	for i:=0;i<20;i++{
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		// From this point lock
		{
			// If the lock is already in use, the calling goroutine
			// blocks until the mutex is available.
			mutex.Lock()
			counter++
			fmt.Println(s,i,"Counter:",counter)
			mutex.Unlock()
		}
	}
	wg.Done()
}

// To check for Race Conditions
// go run -race main.go