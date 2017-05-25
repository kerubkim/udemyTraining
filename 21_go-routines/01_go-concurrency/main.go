package main

import "fmt"

// Concurrency is the composition of independently executing processes,
// Parallelism is the simultaneous execution of (possibly related) computations.
// Concurrency is about dealing with lots of things at once.
// Parallelism is about doing lots of things at one
// Concurrency doing many things, but only one at a time "Multitasking"
// Parallelism doing many things at the same time



// Adding " go " adds a thread
// currently 3 threads main-foo-bar
// runs but not visible, need to add wait group and a pause
func main(){
	go foo()
	go bar()
}

func foo(){
	for i:=0;i<15;i++{
		fmt.Println(i)
	}
}

func bar(){
	for i:=0;i<15;i++{
		fmt.Println(i)
	}
}
