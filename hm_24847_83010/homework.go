package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	totalWorkers = 10
	totalEvents  = 10
)

type Event struct {
	data        interface{}
	closeSignal bool
}

type EventBroadcaster struct {
	cond          *sync.Cond
	event         Event
	wgBroadcaster *sync.WaitGroup
}

func NewEventBroadcaster() *EventBroadcaster {
	return &EventBroadcaster{
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

func (b *EventBroadcaster) Publish(event Event) {
	b.cond.L.Lock()
	defer b.cond.L.Unlock()

	b.event = event

	b.cond.Broadcast()
}

func main() {
	eventBroadcaster := NewEventBroadcaster()
	var (
		wgReady sync.WaitGroup
		wgSync  sync.WaitGroup
	)

	wgReady.Add(totalWorkers)

	// Обрабатываем события
	for i := 0; i < totalWorkers; i++ {
		wgSync.Add(1)
		go Worker(eventBroadcaster, &wgSync, &wgReady)
	}

	// Подождем пока запустятся все горутины
	wgReady.Wait()

	// Бродкастим события
	for i := 0; i < totalEvents; i++ {
		event := Event{
			data:        fmt.Sprintf("event_%d", i),
			closeSignal: false,
		}
		eventBroadcaster.Publish(event)
		time.Sleep(1 * time.Second)
	}

	// Уведомляем что события кончились
	lastEvent := Event{
		data:        nil,
		closeSignal: true,
	}
	eventBroadcaster.Publish(lastEvent)

	wgSync.Wait()
}

func Worker(b *EventBroadcaster, wgSync *sync.WaitGroup, wgReady *sync.WaitGroup) {
	wgReady.Done()
	defer wgSync.Done()

	for {
		b.cond.L.Lock()
		b.cond.Wait()

		if b.event.closeSignal {
			b.cond.L.Unlock()
			return
		}

		event := b.event

		fmt.Printf("Event: %s \n", event.data)

		b.cond.L.Unlock()
	}
}
