package main

func sort(l, r []int) []int {
	result := make([]int, len(l)+len(r))
	var i, j, k int

	for i < len(l) && j < len(r) {
		if r[j] < l[i] {
			result[k] = r[j]
			j++
		} else {
			result[k] = l[i]
			i++
		}
		k++
	}

	for i < len(l) {
		result[k] = l[i]
		i++
		k++
	}
	for j < len(r) {
		result[k] = r[j]
		j++
		k++
	}

	return result
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	left := mergeSort(arr[:len(arr)/2])
	right := mergeSort(arr[len(arr)/2:])

	return sort(left, right)
}

func findWinners(matches [][]int) [][]int {
	losses := make(map[int]int)
	winners := []int{}
	oneLoss := []int{}
	for _, i := range matches {
		losses[i[0]] += 0
		losses[i[1]] += 1
	}
	for k, v := range losses {
		if v == 0 {
			winners = append(winners, k)
		} else if v == 1 {
			oneLoss = append(oneLoss, k)
		}
	}

	return [][]int{mergeSort(winners), mergeSort(oneLoss)}
}
