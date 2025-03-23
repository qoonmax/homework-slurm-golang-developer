Напишите worker pool на основе каналов.

Одна горутина создает “задачи”, набор горутин выполняет эти “задачи”.

В качестве задачи выполните расчет md5 для переданной строки.

```go
package homework

import (
	"crypto/md5"
	"encoding/hex"
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
	// Ваш код
}
```