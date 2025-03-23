package homework

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func SendRequest(url string) (int, error) {
	resp, err := http.Get(url)

	if err != nil {
		return 0, fmt.Errorf("request failed: %v", err)
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			log.Println("failed to close response body:", err)
		}
	}(resp.Body)

	return resp.StatusCode, nil
}
