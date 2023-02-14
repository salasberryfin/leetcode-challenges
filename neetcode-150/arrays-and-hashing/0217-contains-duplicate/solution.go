package main

func containsDuplicate(nums []int) bool {
	var result bool
	occurrences := map[int]int{}
	for _, num := range nums {
		if _, ok := occurrences[num]; ok {
			return true
		}
		occurrences[num] = occurrences[num] + 1
	}

	return result
}
