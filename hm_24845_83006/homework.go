package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaphore chan struct{}

func NewSemaphore(n int) Semaphore {
	return make(Semaphore, n)
}

func main() {
	const totalRequests = 100

	var wg sync.WaitGroup
	semaphore := NewSemaphore(2)

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			semaphore.Acquire(1)
			defer semaphore.Release(1)

			fmt.Printf("Запрос %d: Работаю\n", i)
			time.Sleep(1 * time.Second)

			fmt.Printf("Запрос %d: Завершил\n", i)
		}(i)
	}

	wg.Wait()
	close(semaphore)
}

func (s Semaphore) Acquire(n int) {
	for i := 0; i < n; i++ {
		s <- struct{}{}
	}
}

func (s Semaphore) Release(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}
