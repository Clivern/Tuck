// Epoch
//
// A common requirement in programs is getting the number of seconds, milliseconds,
// or nanoseconds since the Unix epoch. Here's how to do it in Go.
//
// Run: go run ./65_epoch
package main

import (
	"fmt"
	"time"
)

func main() {
	// Use time.Now with Unix, UnixMilli or UnixNano to get elapsed time since the
	// Unix epoch in seconds, milliseconds or nanoseconds, respectively. We use a
	// fixed time here so the output is the same on every run.
	t := time.Date(2026, time.October, 18, 12, 0, 0, 0, time.UTC)
	fmt.Println(t)

	fmt.Println(t.Unix())
	fmt.Println(t.UnixMilli())
	fmt.Println(t.UnixNano())

	// You can also convert integer seconds or nanoseconds since the epoch into the
	// corresponding time.
	fmt.Println(time.Unix(t.Unix(), 0).UTC())
	fmt.Println(time.Unix(0, t.UnixNano()).UTC())
	fmt.Println(time.UnixMilli(t.UnixMilli()).UTC())

	// The epoch itself.
	fmt.Println(time.Unix(0, 0).UTC())

	// The current time works the same way; it is just different on every run.
	fmt.Println("now is after the fixed time:", time.Now().Unix() > t.Unix())
}
