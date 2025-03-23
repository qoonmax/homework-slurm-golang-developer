package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sync"
)

// Worker represents a worker that can process tasks.
type Worker struct {
	// Channel to receive tasks.
	tasks <-chan string
	// WaitGroup to signal when the worker is done.
	wg *sync.WaitGroup
	// Channel to write results
	out chan string
}

func main() {
	const totalWorkers = 10
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

	for w := 0; w < totalWorkers; w++ {
		wg.Add(1)
		worker := NewWorker(tasks, &wg, results)
		go worker.Run()
	}

	wg.Wait()
	close(results)

	for r := range results {
		fmt.Println(r)
	}
}

// NewWorker creates a new worker.
func NewWorker(tasks <-chan string, wg *sync.WaitGroup, out chan string) *Worker {
	return &Worker{
		tasks: tasks,
		wg:    wg,
		out:   out,
	}
}

// Run starts the worker.
func (w *Worker) Run() {
	defer w.wg.Done()

	for t := range w.tasks {
		hash := md5.Sum([]byte(t))
		hashStr := hex.EncodeToString(hash[:]) // [:] create slice from array
		w.out <- hashStr
	}
}
