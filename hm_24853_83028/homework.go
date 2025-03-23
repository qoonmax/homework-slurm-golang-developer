package homework

import (
	"bufio"
	"io"
)

// x2 по памяти, но нет утечек как в defer, или лишних перемещений как с append([]string{line}, result...)
func reverseReader(reader io.Reader) (result []string, err error) {
	scanner := bufio.NewScanner(reader)

	// Сначала собираем строки в срез
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line) // Добавляем в конец среза
	}

	// Проверка на ошибки при сканировании
	if err = scanner.Err(); err != nil {
		return nil, err
	}

	// Разворачиваем срез в обратном порядке
	for i := len(lines) - 1; i >= 0; i-- {
		result = append(result, lines[i])
	}

	return result, nil
}
