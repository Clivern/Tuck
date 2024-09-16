// Context
//
// A context.Context carries deadlines, cancellation signals, and other
// request-scoped values across API boundaries and goroutines. Functions that do I/O
// or long work take a ctx as their first parameter and stop when ctx.Done() is
// closed.
//
// Run: go run ./51_context
package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// slowOperation does work for d, but gives up as soon as the context is done.
func slowOperation(ctx context.Context, d time.Duration) error {
	select {
	case <-time.After(d):
		return nil
	case <-ctx.Done():
		// ctx.Err() says why: context.Canceled or context.DeadlineExceeded.
		return ctx.Err()
	}
}

// A private key type avoids collisions between packages storing values.
type ctxKey string

func handle(ctx context.Context) {
	fmt.Println("handling request", ctx.Value(ctxKey("requestID")))
}

func main() {
	// context.Background is the root of every context tree.
	ctx := context.Background()

	// WithTimeout returns a context that is cancelled after the timeout. Always call
	// cancel to release its resources, usually with defer.
	tctx, cancel := context.WithTimeout(ctx, 50*time.Millisecond)
	defer cancel()

	err := slowOperation(tctx, 200*time.Millisecond)
	fmt.Println("slow:", err, errors.Is(err, context.DeadlineExceeded))

	err = slowOperation(ctx, 10*time.Millisecond)
	fmt.Println("fast:", err)

	// WithCancel lets us cancel by hand, here from another goroutine. Cancelling a
	// parent also cancels all of its children.
	cctx, cancel2 := context.WithCancel(ctx)
	go func() {
		time.Sleep(20 * time.Millisecond)
		cancel2()
	}()
	fmt.Println("cancelled:", slowOperation(cctx, time.Second))

	// WithCancelCause records why the context was cancelled.
	cause, cancel3 := context.WithCancelCause(ctx)
	cancel3(errors.New("user logged out"))
	fmt.Println("cause:", context.Cause(cause))

	// WithValue attaches request-scoped data. Use it for things like request IDs,
	// not for passing optional parameters.
	vctx := context.WithValue(ctx, ctxKey("requestID"), "req-42")
	handle(vctx)

	// AfterFunc runs a function when the context is done.
	actx, cancel4 := context.WithCancel(ctx)
	stopped := make(chan struct{})
	context.AfterFunc(actx, func() {
		fmt.Println("cleanup after cancel")
		close(stopped)
	})
	cancel4()
	<-stopped
}
