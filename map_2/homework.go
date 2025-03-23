package homework

func mapKeyIntersect(m1 map[int]struct{}, m2 map[int]struct{}) []int {
	result := make([]int, 0, len(m1))
	for k, _ := range m1 {
		if _, ok := m2[k]; ok {
			result = append(result, k)
		}
	}

	return result
}
