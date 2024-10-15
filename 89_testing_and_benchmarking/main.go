// Testing and Benchmarking
//
// Unit testing is an important part of writing principled Go programs. The testing
// package provides the tools we need to write unit tests and the go test command
// runs tests. The tests for this example live next to it in main_test.go.
//
// Run: go run ./89_testing_and_benchmarking
// Run: go test -count=1 -v ./89_testing_and_benchmarking
// Run: go test -count=1 -run=XXX -bench=. -benchtime=100000x ./89_testing_and_benchmarking
// Run: go test -count=1 -cover ./89_testing_and_benchmarking
// Run: go test -count=1 -race -run=TestIntMinBasic ./89_testing_and_benchmarking
package main

import "fmt"

// We'll be testing this simple implementation of an integer minimum. Typically, the
// code we're testing would be in a source file named something like intutils.go,
// and the test file for it would then be named intutils_test.go.
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(IntMin(2, -2))
}
