package main

func findErrorNums(nums []int) []int {
	occurrences := make(map[int]int)
	for _, v := range nums {
		occurrences[v]++
	}
	var missing, duplicate int
	for i := 1; i < len(nums)+1; i++ {
		if occurrences[i] == 0 {
			// this means missing value
			missing = i
			if duplicate != 0 {
				return []int{duplicate, missing}
			}
		} else if occurrences[i] == 2 {
			// this means duplicate value
			duplicate = i
			if missing != 0 {
				return []int{duplicate, missing}
			}
		}
	}

	return []int{duplicate, missing}
}
