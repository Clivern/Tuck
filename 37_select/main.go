// Select
//
// Go's select lets you wait on multiple channel operations. Combining goroutines and
// channels with select is a powerful feature of Go.
//
// Run: go run ./37_select
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// For our example we'll select across two channels.
	c1 := make(chan string)
	c2 := make(chan string)

	// Each channel will receive a value after some amount of time, to simulate
	// e.g. blocking RPC operations executing in concurrent goroutines.
	go func() {
		time.Sleep(100 * time.Millisecond)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(200 * time.Millisecond)
		c2 <- "two"
	}()

	// We'll use select to await both of these values simultaneously, printing each
	// one as it arrives.
	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	// We receive the values "one" and then "two" as expected. Note that the total
	// execution time is only about 200ms, since both sleeps execute concurrently.
	fmt.Println("took about", time.Since(start).Round(100*time.Millisecond))
}
