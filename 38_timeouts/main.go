// Timeouts
//
// Timeouts are important for programs that connect to external resources or that
// otherwise need to bound execution time. Implementing timeouts in Go is easy and
// elegant thanks to channels and select.
//
// Run: go run ./38_timeouts
package main

import (
	"fmt"
	"time"
)

func main() {
	// For our example, suppose we're executing an external call that returns its
	// result on a channel c1 after 200ms. Note that the channel is buffered, so the
	// send in the goroutine is nonblocking. This is a common pattern to prevent
	// goroutine leaks in case the channel is never read.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(200 * time.Millisecond)
		c1 <- "result 1"
	}()

	// Here's the select implementing a timeout. res := <-c1 awaits the result and
	// <-time.After awaits a value to be sent after the timeout of 100ms. Since
	// select proceeds with the first receive that's ready, we'll take the timeout
	// case if the operation takes more than the allowed 100ms.
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timeout 1")
	}

	// If we allow a longer timeout of 300ms, then the receive from c2 will succeed
	// and we'll print the result.
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(200 * time.Millisecond)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(300 * time.Millisecond):
		fmt.Println("timeout 2")
	}
}
