package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func sendGetRequest(ctx context.Context) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://www.google.com", nil)
	if err != nil {
		fmt.Println("Ошибка создания запроса:", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Статус-код ответа:", resp.Status)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sendGetRequest(ctx)
}
