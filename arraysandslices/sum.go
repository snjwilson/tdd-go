package arraysandslices

// Sum returns the sum of all the elements in an array
func Sum(arr []int) int {
	var result int
	for _, num := range arr {
		result += num
	}
	return result
}

// SumAll returns a slice of ints where each element is the sum of the corresponding slice passed
func SumAll(collections ...[]int) []int {
	var final []int
	for _,arr := range collections {
		final = append(final, Sum(arr))
	}
	return final
}

// SumAllTrail returns a slice of ints where each element is the sum of the trailing elements except first element of the  corresponding slice passed
func SumAllTrail(collections ...[]int) []int {
	var final []int
	for _,arr := range collections {
		arr = arr[1:]
		final = append(final, Sum(arr))
	}
	return final
}