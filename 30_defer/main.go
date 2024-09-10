// Defer
//
// defer ensures that a function call is performed later in a program's execution,
// usually for purposes of cleanup. It is often used where ensure and finally would
// be used in other languages.
//
// Run: go run ./30_defer
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Suppose we wanted to create a file, write to it, and then close it when we're
// done. Here's how we could do that with defer.
func main() {
	path := filepath.Join(os.TempDir(), "defer.txt")

	// Immediately after getting a file object with createFile, we defer the closing
	// of that file with closeFile. This will be executed at the end of the enclosing
	// function (main), after writeFile has finished.
	f := createFile(path)
	defer closeFile(f)
	writeFile(f)

	// Deferred calls run in last-in, first-out order, and their arguments are
	// evaluated when the defer statement runs.
	for i := range 3 {
		defer fmt.Println("deferred", i)
	}

	// A deferred closure can read and change named results.
	fmt.Println("double:", double(21))
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

// It's important to check for errors when closing a file, even in a deferred
// function.
func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	os.Remove(f.Name())

	if err != nil {
		panic(err)
	}
}

func double(n int) (result int) {
	defer func() { result *= 2 }()
	return n
}
