package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

//go:generate mockgen -source $GOFILE -destination ./mock.go -package ${GOPACKAGE}

// HttpClient Интерфейс для мокирования
type HttpClient interface {
	Get(url string) (*http.Response, error)
}

// NewHTTPClient Создание HTTP-клиента с таймаутом
func NewHTTPClient() HttpClient {
	return &http.Client{Timeout: 10 * time.Second}
}

// Явные ошибки
var (
	ErrUnexpectedStatusCode = errors.New("unexpected status code")
	ErrRequestFailed        = errors.New("request failed")
)

func main() {
	client := NewHTTPClient()

	success, err := sendRequest(client, "https://www.google.com")
	if err != nil {
		if errors.Is(err, ErrUnexpectedStatusCode) {
			fmt.Printf("Ошибка: получен неожиданный статус код: %s\n", err)
		} else if errors.Is(err, ErrRequestFailed) {
			fmt.Printf("Ошибка при выполнении запроса: %s\n", err)
		} else {
			fmt.Printf("Неизвестная ошибка: %s\n", err)
		}
		return
	}

	fmt.Println("Успех:", success)
}

// Выполнение HTTP-запроса
func sendRequest(client HttpClient, url string) (bool, error) {
	resp, err := client.Get(url)
	if err != nil {
		return false, fmt.Errorf("%w: %v", ErrRequestFailed, err)
	}
	if resp != nil {
		defer func() {
			if err = resp.Body.Close(); err != nil {
				fmt.Printf("Ошибка при закрытии тела ответа: %s\n", err)
			}
		}()
	}

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}

	// Возвращаем ошибку с кодом статуса
	return false, fmt.Errorf("%w: %d", ErrUnexpectedStatusCode, resp.StatusCode)
}
