// Switch
//
// switch expresses conditionals across many branches. Cases run top to bottom, the
// first match wins, and there is no fall through by default.
//
// Run: go run ./07_switch
package main

import (
	"fmt"
	"time"
)

func main() {
	// A basic switch.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Commas separate several values in one case. default runs when nothing
	// matches.
	day := time.Date(2026, time.October, 18, 0, 0, 0, 0, time.UTC).Weekday()
	switch day {
	case time.Saturday, time.Sunday:
		fmt.Println(day, "is the weekend")
	default:
		fmt.Println(day, "is a weekday")
	}

	// A switch without an expression is a cleaner way to write if/else chains. Case
	// expressions do not have to be constants.
	hour := 15
	switch {
	case hour < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// fallthrough continues into the next case without testing it.
	switch level := 2; level {
	case 2:
		fmt.Println("level 2")
		fallthrough
	case 1:
		fmt.Println("level 1 too")
	}

	// A type switch compares types instead of values. Here it finds out the type of
	// an interface value; t has the matching type inside each case.
	whatAmI := func(i any) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int, doubled:", t*2)
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
