<p align="center">
    <img src="assets/logo.svg" width="180" alt="Tuck logo">
    <h3 align="center">Tuck</h3>
    <p align="center">A hands-on introduction to Go, one annotated example at a time.</p>
</p>


## What's inside

Tuck is a learning playground in the spirit of [Go by Example](https://gobyexample.com/):
91 numbered Go programs, one topic each. They start with the language basics and work
through types, generics, errors, and concurrency, then cover a wide part of the standard
library: text, encoding, time, files and I/O, the command line, networking, processes,
testing, and reflection.

Every example lives in its own folder as a complete `main.go`. The comments explain the
code right below them, and the header lists the commands that run it. Code the compiler
rejects is shown commented out, followed by the real error it produces. Each example runs
on its own:

```zsh
$ go run ./01_hello_world
$ go test -v ./89_testing_and_benchmarking
$ printf 'hello\nfilter\n' | go run ./72_line_filters
```

The examples need Go 1.27 or newer.

## Website

The examples are also published as a website on GitHub Pages. Each page shows the
comments next to the code they describe and the real output of the example's commands.
The Pages workflow checks formatting, vets, tests, rebuilds, and deploys it on every push
to `main`. To preview it locally:

```zsh
$ python3 site/build.py            # runs every example and writes docs/
$ python3 -m http.server -d docs   # preview at http://localhost:8000
```

## References

- [A Tour of Go](https://go.dev/tour/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [The Go Programming Language Specification](https://go.dev/ref/spec)
- [Standard library docs](https://pkg.go.dev/std)
