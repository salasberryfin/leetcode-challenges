package main

func uniqueOccurrences(arr []int) bool {
	occurrences := make(map[int]int)
	uniqueness := make(map[int]bool)
	for _, v := range arr {
		occurrences[v]++
	}
	for _, v := range occurrences {
		uniqueness[v] = true
	}

	return len(occurrences) == len(uniqueness)
}
