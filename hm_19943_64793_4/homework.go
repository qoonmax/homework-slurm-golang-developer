package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sync"
	"time"
)

// Worker represents a worker that can process tasks.
type Worker struct {
	// Channel to receive tasks.
	tasks <-chan string
	// WaitGroup to signal when the worker is done.
	wg *sync.WaitGroup
	// Channel to write results
	out chan string
	// Context for cancel
	ctx context.Context
}

func main() {
	const totalWorkers = 2
	const totalTasks = 10

	tasks := make(chan string, totalTasks)
	results := make(chan string, totalTasks)

	var wg sync.WaitGroup

	go func() {
		for i := 0; i < totalTasks; i++ {
			tasks <- "task"
		}
		close(tasks)
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for w := 0; w < totalWorkers; w++ {
		wg.Add(1)
		worker := NewWorker(tasks, &wg, results, ctx)
		go worker.Run()
	}

	wg.Wait()
	close(results)

	for r := range results {
		fmt.Println(r)
	}
}

// NewWorker creates a new worker.
func NewWorker(tasks <-chan string, wg *sync.WaitGroup, out chan string, ctx context.Context) *Worker {
	return &Worker{
		tasks: tasks,
		wg:    wg,
		out:   out,
		ctx:   ctx,
	}
}

// Run starts the worker.
func (w *Worker) Run() {
	defer w.wg.Done()

	for {
		select {
		case <-w.ctx.Done(): // Если контекст отменен, выходим
			fmt.Println("Worker shutting down...")
			return
		case t, ok := <-w.tasks:
			if !ok {
				return // Если канал закрыт, выходим
			}

			// Обрабатываем задачу
			hash := md5.Sum([]byte(t))
			hashStr := hex.EncodeToString(hash[:])

			fmt.Println("Processing:", t)

			w.out <- hashStr
			time.Sleep(4 * time.Second) // Симуляция работы
		}
	}
}
