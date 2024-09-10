// Recover
//
// Go makes it possible to recover from a panic, by using the recover built-in
// function. A recover can stop a panic from aborting the program and let it
// continue with execution instead.
//
// An example of where this can be useful: a server wouldn't want to crash if one of
// the client connections exhibits a critical error. Instead, the server would want
// to close that connection and continue serving other clients. In fact, this is
// what Go's net/http does by default for HTTP servers.
//
// Run: go run ./31_recover
package main

import (
	"fmt"
	"strconv"
)

// This function panics.
func mayPanic() {
	panic("a problem")
}

// recover turns a panic back into an error, here through a named result.
func parse(s string) (n int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("parse %q: %v", s, r)
		}
	}()
	n, perr := strconv.Atoi(s)
	if perr != nil {
		panic("not a number")
	}
	return n, nil
}

func main() {
	// recover must be called within a deferred function. When the enclosing
	// function panics, the defer activates and a recover call within it catches
	// the panic.
	defer func() {
		// The return value of recover is the error raised in the call to panic.
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	fmt.Println(parse("12"))
	fmt.Println(parse("twelve"))

	mayPanic()

	// This code will not run, because mayPanic panics. The execution of main stops
	// at the point of the panic and resumes in the deferred closure.
	fmt.Println("After mayPanic()")
}
