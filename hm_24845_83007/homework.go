package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	tasks        <-chan string
	wg           *sync.WaitGroup
	cond         *sync.Cond
	healthStatus *bool
}

func NewWorker(
	tasks <-chan string,
	wg *sync.WaitGroup,
	cond *sync.Cond,
	healthStatus *bool,
) *Worker {
	return &Worker{
		tasks:        tasks,
		wg:           wg,
		cond:         cond,
		healthStatus: healthStatus,
	}
}

type HealthChecker struct {
	cond         *sync.Cond
	healthStatus *bool
}

func NewHealthChecker(cond *sync.Cond, healthStatus *bool) *HealthChecker {
	return &HealthChecker{
		cond:         cond,
		healthStatus: healthStatus,
	}
}

func main() {
	const totalWorker = 10
	const totalTasks = 100

	var (
		mu           sync.Mutex
		cond         = sync.NewCond(&mu)
		healthStatus = false
		wg           sync.WaitGroup
	)

	tasks := make(chan string, totalTasks)

	// run task	provider
	go func() {
		for i := 0; i < totalTasks; i++ {
			tasks <- fmt.Sprintf("task_%d", i)
		}
		close(tasks)
	}()

	// run health check
	healthChecker := NewHealthChecker(cond, &healthStatus)
	go healthChecker.Run()

	// run workers
	for i := 0; i < totalWorker; i++ {
		wg.Add(1)
		worker := NewWorker(tasks, &wg, cond, &healthStatus)
		go worker.Run()
	}

	wg.Wait()
}

func (w *Worker) Run() {
	defer w.wg.Done()

	for {
		w.cond.L.Lock()
		for !*w.healthStatus {
			w.cond.Wait()
		}

		task, ok := <-w.tasks
		if !ok {
			w.cond.L.Unlock()
			return
		}

		w.cond.L.Unlock()

		fmt.Printf("Working: %s \n", task)
	}

}

func (hc *HealthChecker) Run() {
	for {
		// Период когда внешний сервис жив
		time.Sleep(5 * time.Microsecond)

		hc.cond.L.Lock()
		*hc.healthStatus = !*hc.healthStatus
		hc.cond.L.Unlock()
		hc.cond.Broadcast()

		// Сделаем задержку в мертвом состоянии сервиса
		if !*hc.healthStatus {
			time.Sleep(5 * time.Second)
		}
	}
}
