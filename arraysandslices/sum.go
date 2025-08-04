package arraysandslices

func Sum(arr []int) int {
	var result int
	for _, num := range arr {
		result += num
	}
	return result
}