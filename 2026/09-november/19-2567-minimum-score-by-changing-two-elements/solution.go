package main

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if right[j] < left[i] {
			result = append(result, right[j])
			j++
		} else {
			result = append(result, left[i])
			i++
		}
		k++
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
		k++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++
		k++
	}

	return result
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := nums[:mid]
	right := nums[mid:]

	return merge(mergeSort(left), mergeSort(right))
}

func minimizeSum(nums []int) int {
	var score int

	// three options are:
	// - change two lowest to next lowest
	// - change two largest to next largest
	// - change lowest and largest to next lowest or next largest
	sorted := mergeSort(nums)
	lenSorted := len(sorted)

	// compare (largest - third lowest), (third largest - lowest), and (second largest - second lowest)
	score = min(sorted[lenSorted-1]-sorted[2], sorted[lenSorted-3]-sorted[0], sorted[lenSorted-2]-sorted[1])

	return score
}
