// HTTP Client
//
// The Go standard library comes with excellent support for HTTP clients and servers
// in the net/http package. In this example we'll use it to issue simple HTTP
// requests. To keep the example self-contained, we first start a tiny test server
// with net/http/httptest and talk to that.
//
// Run: go run ./82_http_client
package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

func main() {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/json":
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, `{"name":"tuck","examples":91}`)
		case "/slow":
			time.Sleep(time.Second)
		case "/echo":
			body, _ := io.ReadAll(r.Body)
			fmt.Fprintf(w, "%s %s", r.Method, body)
		default:
			fmt.Fprint(w, "line one\nline two\nline three\n")
		}
	}))
	defer srv.Close()

	// Issue an HTTP GET request to a server. http.Get is a convenient shortcut
	// around creating an http.Client object and calling its Get method; it uses the
	// http.DefaultClient object which has useful default settings.
	resp, err := http.Get(srv.URL + "/")
	if err != nil {
		panic(err)
	}

	// Always close the body when you're done with it.
	defer resp.Body.Close()

	// Print the HTTP response status.
	fmt.Println("Response status:", resp.Status)

	// Print the first 2 lines of the response body.
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 2; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Decode a JSON response straight from the body.
	resp2, err := http.Get(srv.URL + "/json")
	if err != nil {
		panic(err)
	}
	defer resp2.Body.Close()
	var info struct {
		Name     string
		Examples int
	}
	json.NewDecoder(resp2.Body).Decode(&info)
	fmt.Println(resp2.Header.Get("Content-Type"), info)

	// For other methods and headers, build an http.Request and send it with a
	// client. A Client with a Timeout bounds the whole request.
	client := &http.Client{Timeout: 5 * time.Second}
	req, _ := http.NewRequest(http.MethodPost, srv.URL+"/echo", strings.NewReader("payload"))
	req.Header.Set("Content-Type", "text/plain")
	resp3, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	body, _ := io.ReadAll(resp3.Body)
	resp3.Body.Close()
	fmt.Println(string(body))

	// A request with a context is cancelled when the context ends, here after
	// 100ms.
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	slow, _ := http.NewRequestWithContext(ctx, http.MethodGet, srv.URL+"/slow", nil)
	_, err = client.Do(slow)
	fmt.Println("slow request failed:", err != nil, ctx.Err())
}
