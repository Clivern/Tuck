// TCP Server
//
// The net package provides the building blocks under net/http: listeners and
// connections. A net.Conn is an io.Reader and an io.Writer, so everything from the
// Readers and Writers example applies. Here is a line-based echo server, with a
// client in the same program.
//
// Run: go run ./84_tcp_server
package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

// handle serves one connection: it reads lines and answers each one in upper case.
func handle(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "quit" {
			fmt.Fprintln(conn, "bye")
			return
		}
		fmt.Fprintln(conn, strings.ToUpper(line))
	}
}

func main() {
	// net.Listen opens a listening socket. Port 0 picks a free port.
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	// Accept connections in a loop, serving each in its own goroutine so one slow
	// client cannot block the others.
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				return // the listener was closed
			}
			go handle(conn)
		}
	}()

	// net.Dial connects a client. DialTimeout bounds the time to connect.
	conn, err := net.DialTimeout("tcp", ln.Addr().String(), time.Second)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Deadlines make reads and writes fail instead of hanging forever.
	conn.SetDeadline(time.Now().Add(5 * time.Second))

	reader := bufio.NewReader(conn)
	for _, msg := range []string{"hello", "tcp in go", "quit"} {
		fmt.Fprintln(conn, msg)
		reply, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Printf("sent %q, got %q\n", msg, strings.TrimSpace(reply))
	}

	// After "quit" the server closes the connection, so the next read sees EOF.
	_, err = reader.ReadString('\n')
	fmt.Println("after quit:", err)

	// net also resolves addresses and splits host and port.
	host, port, _ := net.SplitHostPort("example.com:443")
	fmt.Println(host, port, net.ParseIP("192.168.1.10").IsPrivate())
}
