<p align="center">
    <img src="assets/logo.svg" width="180" alt="Tuck logo">
    <h3 align="center">Tuck</h3>
    <p align="center">A hands-on introduction to Go, one annotated example at a time.</p>
</p>


## What's inside

91 small Go programs, one topic each, in the spirit of [Go by Example](https://gobyexample.com/).
They go from the basics to generics, concurrency, and the standard library.

Each folder holds a commented `main.go` you can run on its own (Go 1.27+):

```zsh
$ go run ./01_hello_world
```

## Website

Read the examples with their output at [tuck.clivern.com](https://tuck.clivern.com).
To build and preview it locally:

```zsh
$ python3 site/build.py
$ python3 -m http.server -d docs
```

## References

- [A Tour of Go](https://go.dev/tour/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [The Go Programming Language Specification](https://go.dev/ref/spec)
- [Standard library docs](https://pkg.go.dev/std)
