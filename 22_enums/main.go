// Enums
//
// Enumerated types are a special case of sum types: a type with a fixed number of
// possible values, each with a distinct name. Go has no enum keyword, but enums are
// simple to build from constants and iota.
//
// Run: go run ./22_enums
package main

import "fmt"

// Our enum type ServerState has an underlying int type.
type ServerState int

// The possible values are constants. iota generates successive values: 0, 1, 2, and
// so on.
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// A String method makes the values print as names. This map does the lookup; for
// many values, the stringer tool can generate the method.
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

// Bit flags use iota with a shift, so each value gets its own bit.
type Perm uint8

const (
	Read Perm = 1 << iota
	Write
	Exec
)

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	// A value of type int cannot be passed where ServerState is expected without a
	// conversion, which gives some compile-time type safety.
	ns2 := transition(ns)
	fmt.Println(ns2)

	p := Read | Write
	fmt.Println("can write:", p&Write != 0, "can exec:", p&Exec != 0)
}

// transition emulates a state transition for a server; it takes the current state
// and returns a new one.
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		// Suppose we check some conditions here to determine the next state.
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
