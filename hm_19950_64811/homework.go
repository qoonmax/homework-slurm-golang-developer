package homework

import (
	"runtime"
	"sync"
)

func FibSlow(n int) int {
	if n <= 1 {
		return n
	}
	return FibSlow(n-1) + FibSlow(n-2)
}

func FibFast(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func SumOfSquaresSlow(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v * v
	}
	return sum
}

func SumOfSquaresFast(slice []int) int {
	var wg sync.WaitGroup
	var mu sync.Mutex
	sum := 0
	parts := runtime.GOMAXPROCS(0) // Количество CPU

	partSize := (len(slice) + parts - 1) / parts

	for i := 0; i < parts; i++ {
		wg.Add(1)
		start := i * partSize
		end := start + partSize
		if end > len(slice) {
			end = len(slice)
		}
		go func(s []int) {
			defer wg.Done()
			partSum := 0
			for _, v := range s {
				partSum += v * v
			}
			mu.Lock()
			sum += partSum
			mu.Unlock()
		}(slice[start:end])
	}
	wg.Wait()
	return sum
}
