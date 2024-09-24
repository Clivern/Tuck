// Command-Line Arguments
//
// Command-line arguments are a common way to parameterize execution of programs. For
// example, go run hello.go uses run and hello.go arguments to the go program.
//
// Run: go build -o .bin/ ./77_command_line_arguments
// Run: ./.bin/77_command_line_arguments a b c d
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// os.Args provides access to raw command-line arguments. Note that the first
	// value in this slice is the path to the program, and os.Args[1:] holds the
	// arguments to the program.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// You can get individual args with normal indexing.
	arg := os.Args[3]

	fmt.Println(filepath.Base(argsWithProg[0]), argsWithProg[1:])
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

	// To experiment with command-line arguments it's best to build a binary with go
	// build first, as above. Next we'll look at more advanced command-line
	// processing with flags.
	fmt.Println("count:", len(argsWithoutProg))
}
