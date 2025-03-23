package main

import (
	"io"
	"log"
	"os"
)

func main() {
	fileOrigin, err := os.Open("./origin.txt")
	if err != nil {
		log.Fatalf("Не удалось открыть файл: %v", err)
	}
	defer func(fileOrigin *os.File) {
		err = fileOrigin.Close()
		if err != nil {
			log.Fatalf("Не удалось закрыть файл: %v", err)
		}
	}(fileOrigin)

	fileCopy, err := os.Create("./copy.txt")
	if err != nil {
		log.Fatalf("Не удалось создать файл: %v", err)
	}
	defer func(fileCopy *os.File) {
		err = fileCopy.Close()
		if err != nil {
			log.Fatalf("Не удалось закрыть файл: %v", err)
		}
	}(fileCopy)

	if _, err = io.Copy(fileCopy, fileOrigin); err != nil {
		log.Fatalf("Не удалось копировать содержимое файла: %v", err)
	}

	// Синхронизация данных на диск (опционально)
	if err = fileCopy.Sync(); err != nil {
		log.Fatalf("Не удалось выгрузить данные на диск: %v", err)
	}
}
