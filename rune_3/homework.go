package homework

func frequentRune(str string) rune {
	r := []rune(str)
	m := make(map[rune]int)

	for _, v := range r {
		m[v]++
	}

	maxCount := 0
	var res rune
	for _, v := range m {
		if v > maxCount {
			maxCount = v
			res = rune(v)
		}
	}

	return res
}
