package main

func zeroFilledSubarray(nums []int) int64 {
	var result int64

	current := int64(0)
	for _, i := range nums {
		if i == 0 {
			current++
		} else {
			current = 0
		}
		result += current
	}

	return result
}
