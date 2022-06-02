package sigreload_test

import (
	"context"
	"fmt"
	"sync/atomic"
	"syscall"
	"testing"
	"time"

	"github.com/cristalhq/sigreload"
)

func TestSignal(t *testing.T) {
	var flag int32
	sigreload.Do(context.Background(), func(ctx context.Context) {
		atomic.AddInt32(&flag, 1)
	})

	doSIGHUP()

	time.Sleep(time.Second)

	got := atomic.LoadInt32(&flag)
	if got != 1 {
		t.Fatalf("got %d want 1", got)
	}
}

func doSIGHUP() {
	err := syscall.Kill(syscall.Getpid(), syscall.SIGHUP)
	if err != nil {
		panic(fmt.Sprintf("cannot send SIGHUP to itself: %s", err))
	}
}
