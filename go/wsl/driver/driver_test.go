package driver

import (
	"context"
	"testing"
	"time"
)

func TestExerciseNewRace(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	caps := map[string]interface{}{
		"binary": "ls",
		"args":   []interface{}{},
	}
	d, err := New(ctx, caps)
	if err != nil {
		t.Fatalf("New: %s", err)
	}
	if d == nil {
		t.Fatalf("nil driver")
	}
	ch := make(chan error, 1)
	go func() {
		ch <- d.Wait()
	}()

	d.Kill()
	select {
	case err := <-ch:
		if err != nil {
			t.Fatalf("Wait: %s", err)
		}
	case <-ctx.Done():
		t.Fatalf("timeout waiting for process to close")
	}
}
