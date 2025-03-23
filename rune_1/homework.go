package homework

func reverse(s string) string {
	r := []rune(s)

	for i, v := range r {
		r[len(r)-1-i] = v
	}

	return string(r)
}
