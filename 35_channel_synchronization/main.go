// Channel Synchronization
//
// We can use channels to synchronize execution across goroutines. Here's an example
// of using a blocking receive to wait for a goroutine to finish. When waiting for
// multiple goroutines to finish, you may prefer to use a WaitGroup.
//
// Run: go run ./35_channel_synchronization
package main

import (
	"fmt"
	"time"
)

// This is the function we'll run in a goroutine. The done channel will be used to
// notify another goroutine that this function's work is done.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("done")

	// Send a value to notify that we're done.
	done <- true
}

func main() {
	// Start a worker goroutine, giving it the channel to notify on.
	done := make(chan bool, 1)
	go worker(done)

	// Block until we receive a notification from the worker on the channel.
	<-done

	// If you removed the <-done line, the program would exit before the worker even
	// started. A channel of struct{} is common when only the signal matters, and
	// closing it signals every waiting receiver at once.
	ready := make(chan struct{})
	go func() {
		fmt.Println("setting up")
		close(ready)
	}()
	<-ready
	fmt.Println("ready")
}
