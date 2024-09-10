// Panic
//
// A panic typically means something went unexpectedly wrong. Mostly we use it to
// fail fast on errors that shouldn't occur during normal operation, or that we
// aren't prepared to handle gracefully.
//
// Run: go run ./29_panic
package main

import (
	"fmt"
	"os"
)

// check panics on any error. A helper like this is fine in small scripts, but real
// programs should return errors instead.
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("about to create a file in a directory that does not exist")

	// A common use of panic is to abort if a function returns an error value that we
	// don't know how to (or want to) handle. Here's an example of panicking if we get
	// an unexpected error when creating a new file. panic also accepts any value,
	// such as panic("a problem").
	_, err := os.Create("/no/such/dir/file")
	check(err)

	// Running this program causes it to panic, print an error message and goroutine
	// traces, and exit with a non-zero status. When the panic fires, the program
	// exits without reaching the rest of the code. Note that unlike some languages
	// which use exceptions for handling of many errors, in Go it is idiomatic to use
	// error-indicating return values wherever possible.
	fmt.Println("never printed")
}
