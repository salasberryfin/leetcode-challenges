package main

func containsDuplicate(nums []int) bool {
	valueMap := map[int]bool{}
	for _, i := range nums {
		if valueMap[i] {
			return true
		}
		valueMap[i] = true
	}

	return false
}
