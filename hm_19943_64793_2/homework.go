package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	defer ctx.Done()

	wg.Add(1)
	go eternalGorutine(ctx, &wg)

	wg.Wait()

	fmt.Println("Done.")
}

func eternalGorutine(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		default:
			fmt.Println("Working...")
			time.Sleep(10 * time.Millisecond)
		}
	}
}
