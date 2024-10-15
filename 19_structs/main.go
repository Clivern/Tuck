// Structs
//
// A struct is a typed collection of fields. Structs group data together to form
// records.
//
// Run: go run ./19_structs
package main

import "fmt"

// This person struct type has name and age fields.
type person struct {
	name string
	age  int
}

// newPerson constructs a new person. Go has no constructors; a function named newX
// is the convention.
func newPerson(name string) *person {
	// It is safe to return a pointer to a local variable: it lives on as long as
	// something refers to it.
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	// This syntax creates a new struct.
	fmt.Println(person{"Bob", 20})

	// You can name the fields when initializing a struct.
	fmt.Println(person{name: "Alice", age: 30})

	// Omitted fields get their zero value.
	fmt.Println(person{name: "Fred"})

	// An & prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40})

	// It's idiomatic to wrap creation in a constructor function.
	fmt.Println(newPerson("Jon"))

	// Access fields with a dot.
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// Dots work with pointers too; they are dereferenced automatically.
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable.
	sp.age = 51
	fmt.Println(sp.age)

	// Assigning a struct copies it, and == compares the fields.
	t := s
	t.name = "Tom"
	fmt.Println(s.name, t.name, s == person{"Sean", 51})

	// %+v prints field names, %#v prints Go syntax.
	fmt.Printf("%+v\n%#v\n", s, s)

	// An anonymous struct type is useful for one-off values, such as table-driven
	// tests.
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

	// A struct can be compared with == when every field can. A slice field
	// makes the whole struct incomparable.
	// fmt.Println(struct{ xs []int }{} == struct{ xs []int }{})
	// error: invalid operation: struct{xs []int}{} == struct{xs []int}{} (struct containing []int cannot be compared)
}
