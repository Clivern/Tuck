// Timers
//
// We often want to execute Go code at some point in the future, or repeatedly at some
// interval. Go's built-in timer and ticker features make both of these tasks easy.
// We'll look first at timers and then at tickers.
//
// Run: go run ./42_timers
package main

import (
	"fmt"
	"time"
)

func main() {
	// Timers represent a single event in the future. You tell the timer how long you
	// want to wait, and it provides a channel that will be notified at that time.
	// This timer will wait 200ms.
	timer1 := time.NewTimer(200 * time.Millisecond)

	// The <-timer1.C blocks on the timer's channel C until it sends a value
	// indicating that the timer fired.
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// If you just wanted to wait, you could have used time.Sleep. One reason a timer
	// may be useful is that you can cancel the timer before it fires. Here's an
	// example of that.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// Give timer2 enough time to fire, if it ever was going to, to show it is in
	// fact stopped.
	time.Sleep(300 * time.Millisecond)

	// time.AfterFunc runs a function in its own goroutine when the timer fires.
	done := make(chan struct{})
	time.AfterFunc(50*time.Millisecond, func() {
		fmt.Println("AfterFunc ran")
		close(done)
	})
	<-done

	// Reset reuses a timer with a new duration.
	t := time.NewTimer(time.Hour)
	t.Reset(10 * time.Millisecond)
	<-t.C
	fmt.Println("reset timer fired")
}
