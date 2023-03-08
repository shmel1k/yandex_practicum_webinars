package scheduler

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestChecker(t *testing.T) {
	data := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		data = append(data, "abacaba"+strconv.Itoa(rand.Intn(100000)))
	}
	m := MyChecker(data)
	done := make(chan struct{})
	go func() {
		m.Wait()
		close(done)
	}()
	select {
	case <-time.After(30 * time.Second):
		t.Fatal("failed to execute task within 30 seconds")
	case <-done:
		return
	}
}
