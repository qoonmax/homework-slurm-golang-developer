package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	filename := fmt.Sprintf("worker_%d.txt", id)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC|os.O_SYNC, 0644)
	if err != nil {
		log.Printf("Worker %d: failed to create file: %v", id, err)
		return
	}
	defer func() {
		file.Close()
		log.Printf("Worker %d: file closed", id)
	}()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Printf("Worker %d: received shutdown signal", id)
			return
		case t := <-ticker.C:
			_, err := file.WriteString(fmt.Sprintf("Worker %d: %s\n", id, t.Format(time.RFC3339)))
			if err != nil {
				log.Printf("Worker %d: failed to write to file: %v", id, err)
			}
		}
	}
}

func main() {
	const gorutineCount = 3

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	for i := 1; i <= gorutineCount; i++ {
		wg.Add(1)
		go worker(ctx, &wg, i)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	sig := <-sigChan
	log.Printf("Received signal: %s", sig)
	cancel()

	wg.Wait()
	log.Println("All workers have shut down gracefully.")
}
