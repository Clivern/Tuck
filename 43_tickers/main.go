// Tickers
//
// Timers are for when you want to do something once in the future; tickers are for
// when you want to do something repeatedly at regular intervals. Here's an example
// of a ticker that ticks periodically until we stop it.
//
// Run: go run ./43_tickers
package main

import (
	"fmt"
	"time"
)

func main() {
	// Tickers use a similar mechanism to timers: a channel that is sent values. Here
	// we'll use the select built-in on the channel to await the values as they
	// arrive every 100ms.
	start := time.Now()
	ticker := time.NewTicker(100 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t.Sub(start).Round(100*time.Millisecond))
			}
		}
	}()

	// Tickers can be stopped like timers. Once a ticker is stopped it won't receive
	// any more values on its channel. We'll stop ours after 350ms.
	time.Sleep(350 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
