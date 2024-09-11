// Channel Buffering
//
// By default channels are unbuffered, meaning that they will only accept sends if
// there is a corresponding receive ready to receive the sent value. Buffered
// channels accept a limited number of values without a corresponding receiver.
//
// Run: go run ./34_channel_buffering
package main

import "fmt"

func main() {
	// Here we make a channel of strings buffering up to 2 values.
	messages := make(chan string, 2)

	// Because this channel is buffered, we can send these values into the channel
	// without a corresponding concurrent receive.
	messages <- "buffered"
	messages <- "channel"

	// len and cap report how many values are waiting and how many fit.
	fmt.Println("len", len(messages), "cap", cap(messages))

	// Later we can receive these two values as usual.
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// A third send without a receiver would block forever. With no other goroutine
	// running, the runtime detects that and stops the program with "fatal error:
	// all goroutines are asleep - deadlock!".
}
