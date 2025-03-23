package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type CircuitBreaker struct {
	mu               sync.Mutex
	failureThreshold int           // Порог ошибок для разрыва цепи
	cooldownPeriod   time.Duration // Время, на которое разрывается цепь
	lastFailureTime  time.Time     // Время последней ошибки
	failureCount     int           // Текущее количество ошибок
}

func NewCircuitBreaker(failureThreshold int, cooldownPeriod time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		failureThreshold: failureThreshold,
		cooldownPeriod:   cooldownPeriod,
	}
}

func (cb *CircuitBreaker) Execute(ctx context.Context, fn func() error) error {
	cb.mu.Lock()

	// Проверяем, не превышен ли порог ошибок
	if cb.failureCount >= cb.failureThreshold {
		// Проверяем, истек ли период охлаждения
		if time.Since(cb.lastFailureTime) < cb.cooldownPeriod {
			cb.mu.Unlock()
			return errors.New("circuit breaker is open")
		}
		// Сбрасываем счетчик ошибок, если период охлаждения истек
		cb.failureCount = 0
	}
	cb.mu.Unlock()

	// Создаем канал для результата выполнения функции
	resultChan := make(chan error, 1)

	// Запускаем функцию в отдельной горутине
	go func() {
		resultChan <- fn()
	}()

	// Ожидаем завершения функции или отмены контекста
	select {
	case err := <-resultChan:
		cb.mu.Lock()
		defer cb.mu.Unlock()

		if err != nil {
			cb.failureCount++
			cb.lastFailureTime = time.Now()
			return fmt.Errorf("request failed: %w", err)
		}

		// Сбрасываем счетчик ошибок при успешном выполнении
		cb.failureCount = 0
		return nil

	case <-ctx.Done():
		// Если контекст отменен, возвращаем ошибку
		return ctx.Err()
	}
}

func main() {
	// Создаем Circuit Breaker с порогом 3 ошибки и периодом охлаждения 2 секунды
	cb := NewCircuitBreaker(3, 2*time.Second)

	// Имитация вызовов
	for i := 0; i < 10; i++ {
		// Создаем контекст с таймаутом 1 секунда
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		err := cb.Execute(ctx, func() error {
			// Имитация ошибки для первых 5 запросов
			if i < 5 {
				time.Sleep(2 * time.Second) // Имитация долгого выполнения
				return errors.New("service error")
			}
			return nil
		})

		if err != nil {
			fmt.Printf("Request %d failed: %v\n", i, err)
		} else {
			fmt.Printf("Request %d succeeded\n", i)
		}

		time.Sleep(500 * time.Millisecond)
	}
}
