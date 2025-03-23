package homework

func toFrequencyMap(s []string) map[string]int {
	m := make(map[string]int)

	for _, v := range s {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}

	return m
}
