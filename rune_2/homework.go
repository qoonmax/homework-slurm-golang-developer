package homework

func removeSpaces(s string) string {
	r := []rune(s)
	res := make([]rune, 0, len(r))

	for _, v := range r {
		if string(v) == " " {
			res = append(res, v)
		}
	}

	return string(res)
}
