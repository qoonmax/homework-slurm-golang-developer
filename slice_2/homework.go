package homework

func removeDuplicates(slice []int) []int {
	m := map[int]*struct{}{}
	var result []int
	for i := 0; i < len(slice); i++ {
		if _, ok := m[slice[i]]; !ok {
			result = append(result, slice[i])
			m[slice[i]] = nil
		}
	}
	return result
}
