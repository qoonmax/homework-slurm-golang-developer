package homework

import "strings"

func frequentWord(str string) string {
	words := strings.Split(str, " ")
	m := make(map[string]int)

	for _, word := range words {
		if _, ok := m[word]; ok {
			m[word]++
		} else {
			m[word] = 1
		}
	}

	maxCount := 0
	result := ""
	for w, c := range m {
		if c > maxCount {
			maxCount = c
			result = w
		}
	}

	return result
}
