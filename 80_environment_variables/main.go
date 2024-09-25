// Environment Variables
//
// Environment variables are a universal mechanism for conveying configuration
// information to Unix programs. Let's look at how to set, get, and list environment
// variables.
//
// Run: BAR=2 go run ./80_environment_variables
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// To set a key/value pair, use os.Setenv. To get a value for a key, use
	// os.Getenv. This will return an empty string if the key isn't present in the
	// environment.
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
	fmt.Println("BAZ:", os.Getenv("BAZ"))

	// LookupEnv tells an empty value apart from a missing one.
	if _, ok := os.LookupEnv("BAZ"); !ok {
		fmt.Println("BAZ is not set")
	}

	// ExpandEnv replaces $VAR and ${VAR} in a string.
	fmt.Println(os.ExpandEnv("foo is $FOO, bar is ${BAR}"))

	// Use os.Environ to list all key/value pairs in the environment. This returns a
	// slice of strings in the form KEY=value. You can strings.Cut them to get the
	// key and value. Here we look for our two keys.
	fmt.Println()
	for _, e := range os.Environ() {
		key, _, _ := strings.Cut(e, "=")
		if key == "FOO" || key == "BAR" {
			fmt.Println(e)
		}
	}

	// Unsetenv removes a variable.
	os.Unsetenv("FOO")
	fmt.Println("FOO after unset:", os.Getenv("FOO") == "")
}
