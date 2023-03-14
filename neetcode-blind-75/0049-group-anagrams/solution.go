package main

func sort(left, right []byte) []byte {
	sorted := []byte{}
	x, y := 0, 0
	for x < len(left) && y < len(right) {
		if right[y] < left[x] {
			sorted = append(sorted, right[y])
			y++
		} else {
			sorted = append(sorted, left[x])
			x++
		}
	}
	for x < len(left) {
		sorted = append(sorted, left[x])
		x++
	}
	for y < len(right) {
		sorted = append(sorted, right[y])
		y++
	}

	return sorted
}

func mergeSort(s []byte) []byte {
	if len(s) <= 1 {
		return s
	}

	mid := len(s) / 2

	return sort(mergeSort(s[:mid]), mergeSort(s[mid:]))
}

func groupAnagrams(strs []string) [][]string {
	positions := map[string][]string{}
	for i, v := range strs {
		sorted := string(mergeSort([]byte(v)))
		positions[sorted] = append(positions[sorted], strs[i])
	}

	result := [][]string{}
	for k := range positions {
		result = append(result, positions[k])
	}

	return result
}
