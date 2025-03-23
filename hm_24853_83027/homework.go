package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./example.txt")
	if err != nil {
		log.Fatalf("Не удалось открыть файл: %v", err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatalf("Не удалось закрыть файл: %v", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		defer func(line string) {
			fmt.Println(line)
		}(line)
	}
}
