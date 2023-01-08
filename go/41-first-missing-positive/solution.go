package main

func mergeSort(left, right []int) []int {
	var result = []int{}

	var i, j int
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			if left[i] > 0 {
				result = append(result, left[i])
			}
			i++
		} else {
			if right[j] > 0 {
				result = append(result, right[j])
			}
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

func sort(sl []int) []int {
	if len(sl) == 1 {
		return sl
	}

	mid := len(sl) / 2
	left := sort(sl[:mid])
	right := sort(sl[mid:])

	return mergeSort(left, right)
}

func firstMissingPositive(nums []int) int {
	sortedNums := sort(nums)

	var previous = sortedNums[0]
	if previous > 1 || previous < 0 {
		return 1
	}
	for i := 0; i < len(sortedNums); i++ {
		if sortedNums[i]-previous > 1 {
			return previous + 1
		}
		previous = sortedNums[i]
	}

	return previous + 1
}
