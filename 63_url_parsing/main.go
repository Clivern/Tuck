// URL Parsing
//
// URLs provide a uniform way to locate resources. Here's how to parse URLs in Go.
//
// Run: go run ./63_url_parsing
package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	// We'll parse this example URL, which includes a scheme, authentication info,
	// host, port, path, query params, and query fragment.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Parse the URL and ensure there are no errors.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// Accessing the scheme is straightforward.
	fmt.Println(u.Scheme)

	// User contains all authentication info; call Username and Password on this for
	// individual values.
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// The Host contains both the hostname and the port, if present. Use
	// SplitHostPort to extract them.
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// Here we extract the path and the fragment after the #.
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// To get query params in a string of k=v format, use RawQuery. You can also parse
	// query params into a map. The parsed query param maps are from strings to
	// slices of strings, so index into [0] if you only want the first value.
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

	// Building URLs: url.Values encodes query strings, and QueryEscape escapes a
	// single value.
	q := url.Values{}
	q.Set("q", "go by example")
	q.Add("tag", "std")
	q.Add("tag", "net/url")
	built := url.URL{Scheme: "https", Host: "example.com", Path: "/search", RawQuery: q.Encode()}
	fmt.Println(built.String())
	fmt.Println(url.QueryEscape("a b&c"), url.PathEscape("a b/c"))

	// ResolveReference resolves a relative link against a base URL.
	base, _ := url.Parse("https://example.com/docs/intro/")
	ref, _ := url.Parse("../api?v=2")
	fmt.Println(base.ResolveReference(ref))
}
