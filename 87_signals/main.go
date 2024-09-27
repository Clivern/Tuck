// Signals
//
// Sometimes we'd like our Go programs to intelligently handle Unix signals. For
// example, we might want a server to gracefully shutdown when it receives a SIGTERM,
// or a command-line tool to stop processing input if it receives a SIGINT. Here's
// how to handle signals in Go with channels.
//
// Run: go run ./87_signals
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Go signal notification works by sending os.Signal values on a channel. We'll
	// create a channel to receive these notifications. Note that this channel should
	// be buffered.
	sigs := make(chan os.Signal, 1)

	// signal.Notify registers the given channel to receive notifications of the
	// specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// We could receive from sigs here in the main function, but let's see how this
	// could also be done in a separate goroutine, to demonstrate a more realistic
	// scenario of graceful shutdown.
	done := make(chan bool, 1)

	go func() {
		// This goroutine executes a blocking receive for signals. When it gets one
		// it'll print it out and then notify the program that it can finish.
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// The program will wait here until it gets the expected signal (as indicated by
	// the goroutine above sending a value on done) and then exit. Normally you would
	// press ctrl-C; here the program sends SIGINT to itself so it can run
	// unattended.
	fmt.Println("awaiting signal")
	syscall.Kill(os.Getpid(), syscall.SIGINT)
	<-done
	fmt.Println("exiting")
	signal.Stop(sigs)

	// signal.NotifyContext is the modern shortcut: the context is cancelled when a
	// signal arrives, which plugs straight into anything that takes a context.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer stop()
	go func() {
		time.Sleep(20 * time.Millisecond)
		syscall.Kill(os.Getpid(), syscall.SIGTERM)
	}()
	<-ctx.Done()
	fmt.Println("context cancelled:", context.Cause(ctx))
}
