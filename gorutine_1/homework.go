package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int64)
	ch2 := make(chan int64)

	inputs := make([]chan int64, 2)
	inputs[0] = ch1
	inputs[1] = ch2

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch1 <- 1
		}

		close(ch1)
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch2 <- 2
		}

		close(ch2)
	}()

	fmt.Println("inputs:", sumChannels(inputs))

	wg.Wait()
}

func sumChannels(inputs []chan int64) int64 {
	res := make(chan int64)
	var sum int64

	var wg sync.WaitGroup

	for _, ch := range inputs {
		wg.Add(1)
		go func(ch chan int64) {
			for {
				value, ok := <-ch

				if !ok {
					break
				}

				res <- value
			}
			defer wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	for v := range res {
		sum += v
	}

	return sum
}
