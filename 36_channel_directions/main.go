// Channel Directions
//
// When using channels as function parameters, you can specify if a channel is meant
// to only send or receive values. This specificity increases the type-safety of the
// program.
//
// Run: go run ./36_channel_directions
package main

import "fmt"

// This ping function only accepts a channel for sending values. It would be a
// compile-time error to try to receive on this channel.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// The pong function accepts one channel for receives (pings) and a second for sends
// (pongs).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	// A bidirectional channel converts to either direction automatically, but not
	// back.
	var recvOnly <-chan string = pongs
	fmt.Printf("%T\n", recvOnly)

	// recvOnly <- "x"
	// error: invalid operation: cannot send to receive-only channel <-chan string recvOnly (variable of type <-chan string)
}
