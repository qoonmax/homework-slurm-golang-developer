package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var (
		wg      sync.WaitGroup
		counter int64
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5000; i++ {
			atomic.AddInt64(&counter, 1)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5000; i++ {
			atomic.AddInt64(&counter, 1)
		}
	}()

	wg.Wait()

	fmt.Println(counter)
}
