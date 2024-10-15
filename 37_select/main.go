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

	// A nil channel is never ready, so this case stays disabled and default runs.
	// Sending or receiving on a nil channel outside select blocks forever.
	// Closing a nil channel panics.
	var quiet chan string
	select {
	case msg := <-quiet:
		fmt.Println("nil receive:", msg)
	default:
		fmt.Println("nil channel was not ready")
	}
	func() {
		defer func() { fmt.Println("close nil panicked:", recover() != nil) }()
		close(quiet)
	}()

	// When more than one case is ready, select chooses one at random. Over many
	// tries both channels are chosen.
	left := make(chan string, 1)
	right := make(chan string, 1)
	left <- "L"
	right <- "R"
	got := map[string]int{}
	for range 64 {
		select {
		case v := <-left:
			got[v]++
			left <- "L"
		case v := <-right:
			got[v]++
			right <- "R"
		}
	}
	fmt.Println("chose both:", got["L"] > 0 && got["R"] > 0)
}
