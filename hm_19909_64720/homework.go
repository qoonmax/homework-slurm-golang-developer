package homework

import "unicode/utf8"

func reverse(str string) string {
	if !utf8.ValidString(str) {
		return str // Если строка невалидна в UTF-8, просто возвращаем её как есть
	}

	res := []rune(str)

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}

	return string(res)
}
