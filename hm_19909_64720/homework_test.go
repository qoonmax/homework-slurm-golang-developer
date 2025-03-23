package homework

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	testCases := []string{
		"string",
		"xy",
		"example",
	}

	for _, tc := range testCases {
		f.Add(tc) // Добавляем исходные данные
	}

	f.Fuzz(func(t *testing.T, input string) {
		reversed := reverse(input)
		doubleReversed := reverse(reversed)

		if input != doubleReversed {
			t.Errorf("Expected %q but got %q", input, doubleReversed)
		}

		// Проверяем, что результат реверса корректно закодирован в UTF-8
		if utf8.ValidString(input) && !utf8.ValidString(reversed) {
			t.Errorf("Reversed string is not valid UTF-8: %q", reversed)
		}
	})
}
