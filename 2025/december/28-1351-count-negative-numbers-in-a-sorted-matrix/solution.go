package main

func sort(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	var i, j, k int

	for i < len(left) && j < len(right) {
		if right[j] < left[i] {
			result[k] = right[j]
			j++
		} else {
			result[k] = left[i]
			i++
		}
		k++
	}

	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}

	return result
}

func mergeSort(sl []int) []int {
	if len(sl) <= 1 {
		return sl
	}

	left := mergeSort(sl[:len(sl)/2])
	right := mergeSort(sl[len(sl)/2:])

	return sort(left, right)
}

func countNegatives(grid [][]int) int {
	result := 0

	single := []int{}
	for _, row := range grid {
		single = append(single, row...)
	}

	sorted := mergeSort(single)
	for _, num := range sorted {
		if num < 0 {
			result++
		} else {
			return result
		}
	}

	return result
}
