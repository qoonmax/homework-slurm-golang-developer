package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCounterWithWaitGroup(t *testing.T) {
	counter := &Counter{}
	const goroutineCount = 10

	RunIncrementGoroutines(counter, goroutineCount)

	require.Equal(t, goroutineCount, counter.Value(), "Counter value should match the number of goroutines")
}

func TestCounterWithEventually(t *testing.T) {
	counter := &Counter{}
	const goroutineCount = 10

	for i := 0; i < goroutineCount; i++ {
		go func() {
			time.Sleep(100 * time.Millisecond) // Имитация работы
			counter.Increment()
		}()
	}

	require.Eventually(t, func() bool {
		return counter.Value() == goroutineCount
	}, 2*time.Second, 10*time.Millisecond, "Counter value should reach expected value")
}
