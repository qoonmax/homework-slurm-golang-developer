package homework

func stringLengthWithoutSpaces(str string) int {
	counter, m := 0, 0
	for _, s := range str {
		if s != ' ' && s != '\t' {
			counter++
		} else {
			if counter > m {
				m = counter
			}
			counter = 0
		}
	}

	return m
}
