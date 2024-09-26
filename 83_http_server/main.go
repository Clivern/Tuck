// HTTP Server
//
// Writing a basic HTTP server is easy using the net/http package. The server below
// routes requests with http.ServeMux, wraps handlers in middleware, and shuts down
// gracefully. To keep the example self-contained, main also acts as the client.
//
// Run: go run ./83_http_server
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

// A fundamental concept in net/http servers is handlers. A handler is an object
// implementing the http.Handler interface. A common way to write a handler is by
// using the http.HandlerFunc adapter on functions with the appropriate signature.
func hello(w http.ResponseWriter, req *http.Request) {
	// Functions serving as handlers take a http.ResponseWriter and a http.Request
	// as arguments. The response writer is used to fill in the HTTP response. Here
	// our simple response is just "hello\n".
	fmt.Fprintf(w, "hello\n")
}

// This handler does something a little more sophisticated by reading all the HTTP
// request headers and echoing one of them into the response body.
func headers(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "User-Agent: %v\n", req.Header.Get("User-Agent"))
}

// Patterns can match a method and capture path segments, read with PathValue.
func getUser(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "user %s\n", req.PathValue("id"))
}

// Middleware is a function that wraps a handler to add behavior, such as logging.
func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s (fast: %t)", r.Method, r.URL.Path, time.Since(start) < time.Second)
	})
}

func main() {
	log.SetFlags(0)

	// We register our handlers on server routes using a ServeMux.
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/headers", headers)
	mux.HandleFunc("GET /users/{id}", getUser)

	// Usually you'd call http.ListenAndServe(":8090", handler) and let it run.
	// Here we listen on a random free port so we can make requests from main.
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	srv := &http.Server{Handler: logging(mux), ReadHeaderTimeout: 5 * time.Second}
	go srv.Serve(ln)
	base := "http://" + ln.Addr().String()

	for _, path := range []string{"/hello", "/headers", "/users/42", "/missing"} {
		req, _ := http.NewRequest("GET", base+path, nil)
		req.Header.Set("User-Agent", "tuck-example")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("%s -> %d %s", path, resp.StatusCode, body)
	}

	// The method in the pattern is enforced: a POST gets 405 Method Not Allowed.
	resp, _ := http.Post(base+"/users/42", "text/plain", nil)
	resp.Body.Close()
	fmt.Println("POST /users/42 ->", resp.Status)

	// Shutdown stops accepting new connections and waits for active requests to
	// finish, up to the context's deadline.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Println("shutdown:", srv.Shutdown(ctx))
}
