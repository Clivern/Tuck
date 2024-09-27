// Exit
//
// Use os.Exit to immediately exit with a given status.
//
// Run: go run ./88_exit
package main

import (
	"fmt"
	"os"
)

func main() {
	// defers will not be run when using os.Exit, so this fmt.Println will never be
	// called.
	defer fmt.Println("!")

	fmt.Println("exiting with status 3")

	// Exit with status 3.
	os.Exit(3)

	// Note that unlike e.g. C, Go does not use an integer return value from main to
	// indicate exit status. If you'd like to exit with a non-zero status you should
	// use os.Exit. When run with go run, the exit is picked up by go and printed as
	// "exit status 3".
}
