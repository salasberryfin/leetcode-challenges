package main

func missingNumber(nums []int) int {
	existing := map[int]bool{}
	for _, i := range nums {
		existing[i] = true
	}

	for i := 0; i <= len(nums); i++ {
		if _, ok := existing[i]; !ok {
			return i
		}
	}

	return 1
}
