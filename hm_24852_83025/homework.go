package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Example struct {
	X int    `json:"x"`
	Y string `json:"y,omitempty"`
}

func main() {
	content, err := os.ReadFile("./example.json")
	if err != nil {
		log.Fatalf("Не удалось прочитать файл: %v", err)
	}

	var example Example

	err = json.Unmarshal(content, &example)
	if err != nil {
		log.Fatalf("Не удалось распарсить файл: %v", err)
	}

	fmt.Println(example)
}
