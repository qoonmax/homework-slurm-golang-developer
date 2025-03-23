package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

const reqCount int = 5

func main() {
	var (
		wg           sync.WaitGroup
		successCount int32
		quorum       = reqCount/2 + 1
	)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	for i := 0; i < reqCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if sendRequest(ctx) {
				atomic.AddInt32(&successCount, 1)
			}
		}()
	}

	wg.Wait()

	if successCount >= int32(quorum) {
		fmt.Println("Кворум достигнут")
	} else {
		fmt.Println("Кворум не достигнут")
	}
}

func sendRequest(ctx context.Context) bool {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://www.google.com/404", nil)
	if err != nil {
		log.Printf("Ошибка создания запроса: %v", err)
		return false
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Ошибка выполнения запроса: %v", err)
		return false
	}
	defer resp.Body.Close()

	// Опционально: проверка статус-кода
	// if resp.StatusCode < 200 || resp.StatusCode >= 300 {
	// 	return false
	// }

	return true
}
