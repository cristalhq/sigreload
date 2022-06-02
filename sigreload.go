package sigreload

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// Do the given function on every SIGHUP.
func Do(ctx context.Context, fn func(ctx context.Context)) {
	ch := NewChan()
	go func() {
		for {
			select {
			case <-ch:
				fn(ctx)
			case <-ctx.Done():
				return
			}
		}
	}()
}

// NewChan returns a channel triggered on every SIGHUP.
func NewChan() <-chan os.Signal {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP)
	return ch
}
