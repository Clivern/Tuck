// Reflection
//
// The reflect package lets a program inspect types and values at run time. It powers
// encoding/json, fmt, and text/template. Most code never needs it directly, but it
// explains how those packages read struct tags and fields.
//
// Run: go run ./91_reflection
package main

import (
	"fmt"
	"reflect"
	"strings"
)

// A struct with tags, like the ones encoding/json reads.
type User struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email"`
	Age   int    `json:"age,omitempty"`
	admin bool
}

func (u User) Greeting() string { return "hi " + u.Name }

// validate checks every field tagged validate:"required" in any struct.
func validate(v any) []string {
	var missing []string
	rv := reflect.ValueOf(v)
	rt := rv.Type()
	for i := range rt.NumField() {
		f := rt.Field(i)
		if f.Tag.Get("validate") == "required" && rv.Field(i).IsZero() {
			missing = append(missing, f.Name)
		}
	}
	return missing
}

func main() {
	u := User{Name: "Ada", Email: "ada@example.com"}

	// TypeOf returns the type, and Kind its basic category.
	t := reflect.TypeOf(u)
	fmt.Println(t, t.Kind(), t.NumField())

	// Walk the fields, reading names, types, and tags.
	for i := range t.NumField() {
		f := t.Field(i)
		fmt.Printf("%-6s %-7s json=%q exported=%t\n", f.Name, f.Type, f.Tag.Get("json"), f.IsExported())
	}

	// ValueOf wraps a value. Read fields by name.
	v := reflect.ValueOf(u)
	fmt.Println(v.FieldByName("Email"))

	// To change a value through reflection, start from a pointer and call Elem.
	pv := reflect.ValueOf(&u).Elem()
	pv.FieldByName("Age").SetInt(36)
	fmt.Println(u.Age, pv.FieldByName("Age").CanSet(), pv.FieldByName("admin").CanSet())

	// Methods can be looked up and called too.
	m := v.MethodByName("Greeting")
	fmt.Println(m.Call(nil)[0])

	// A generic validator built on tags.
	fmt.Println("missing:", validate(User{Email: "x@y.z"}))
	fmt.Println("missing:", validate(u))

	// reflect.DeepEqual compares values deeply, including slices and maps.
	a := map[string][]int{"x": {1, 2}}
	b := map[string][]int{"x": {1, 2}}
	fmt.Println(reflect.DeepEqual(a, b))

	// Kind lets code handle whole families of types at once.
	for _, x := range []any{1, "s", []int{}, map[int]int{}, &u, strings.ToUpper} {
		fmt.Print(reflect.TypeOf(x).Kind(), " ")
	}
	fmt.Println()
}
