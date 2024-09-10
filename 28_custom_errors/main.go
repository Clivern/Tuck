// Custom Errors
//
// It's possible to use custom types as errors by implementing the Error() method on
// them. A custom type can carry details that the caller reads back with errors.As.
//
// Run: go run ./28_custom_errors
package main

import (
	"errors"
	"fmt"
)

// A custom error type usually has the suffix "Error".
type argError struct {
	arg     int
	message string
}

// Adding this Error method makes argError implement the error interface.
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {
		// Return our custom error.
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

// An error type can wrap a cause by adding an Unwrap method, so errors.Is and
// errors.As can look inside it.
type queryError struct {
	query string
	err   error
}

func (e *queryError) Error() string { return "query " + e.query + ": " + e.err.Error() }
func (e *queryError) Unwrap() error { return e.err }

var errTimeout = errors.New("timeout")

func main() {
	// errors.As is a more advanced version of errors.Is. It checks that a given
	// error, or any error in its chain, matches a specific error type and converts
	// to a value of that type, returning true. If there's no match, it returns
	// false.
	_, err := f(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}

	// errors.AsType is the generic form: it returns the typed error directly.
	if ae, ok := errors.AsType[*argError](err); ok {
		fmt.Println("AsType:", ae.arg)
	}

	// With Unwrap in place, the cause is reachable through any number of layers.
	err = fmt.Errorf("list users: %w", &queryError{"SELECT *", errTimeout})
	fmt.Println(err)
	fmt.Println("timeout:", errors.Is(err, errTimeout))
}
