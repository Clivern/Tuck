// Packages and Modules
//
// Go code is organized in packages, and packages in modules. A package is a
// directory of .go files that share a package clause. A module is a tree of
// packages with a go.mod file at its root, which names the module path and lists
// its dependencies. This repository is the module github.com/clivern/tuck.
//
// Run: go run ./90_packages_and_modules
// Run: go list ./90_packages_and_modules/...
// Run: head -3 go.mod
package main

// A package is imported by its full path: the module path plus the directory. The
// geometry package lives in the geometry/ folder next to this file.
import (
	"fmt"

	"github.com/clivern/tuck/90_packages_and_modules/geometry"
)

func main() {
	// Exported names are used through the package name.
	c := geometry.Circle{Radius: 2}
	fmt.Println(c.Area(), c.Diameter())

	// Package-level variables are initialized, and init functions run, before main.
	fmt.Println(geometry.Default)

	// Lower-case names are private to their package.
	// geometry.round(1.234)
	// error: name round not exported by package geometry

	// Third-party packages are added with go get, which records them in go.mod and
	// go.sum, for example: go get golang.org/x/sync/errgroup. go mod tidy removes
	// dependencies you no longer use.
	fmt.Println("see the geometry folder for the package source")
}
