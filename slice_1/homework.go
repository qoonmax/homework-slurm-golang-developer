package homework

func reverse(slice []int) []int {
	slice2 := make([]int, len(slice))
	for i := len(slice) - 1; i >= 0; i-- {
		slice2[i] = slice[len(slice2)- 1 - i]
	}
	return slice2
}