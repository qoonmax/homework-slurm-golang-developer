package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	input := make(chan string)
	step1Out := make(chan string)
	step2Out := make(chan string)
	step3Out := step3(step2Out)

	go step1(input, step1Out)
	go step2(step1Out, step2Out)

	// Отправляем данные в pipeline
	go func() {
		input <- "  привет, мир. как дела?  "
		input <- "  это тест.  проверка. "
		input <- "Это текст с    лишними пробелами"
		close(input)
	}()

	// Читаем результат из последнего шага
	for result := range step3Out {
		fmt.Println(result)
	}
}

func step1(in <-chan string, out chan<- string) {
	defer close(out)
	for str := range in {
		str = strings.TrimSpace(str)
		for strings.Contains(str, "  ") {
			str = strings.ReplaceAll(str, "  ", " ")
		}
		out <- str
	}
}

func step2(in <-chan string, out chan<- string) {
	defer close(out)
	for str := range in {
		strSlice := strings.Split(str, ".")
		for _, s := range strSlice {
			s = strings.TrimSpace(s)
			if s != "" {
				out <- s
			}
		}
	}
}

// Обратите внимание, что step3 должен вернуть канал, в который будет записывать.
// Это значит, что внутри функции нужно запустить отдельную горутину, читающую in.
func step3(in <-chan string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		for str := range in { // Читаем из канала с обработкой закрытия
			runes := []rune(str) // Преобразуем в runes для поддержки Unicode
			if len(runes) > 0 {
				runes[0] = unicode.ToUpper(runes[0])
			}
			out <- string(runes)
		}
	}()

	return out
}
