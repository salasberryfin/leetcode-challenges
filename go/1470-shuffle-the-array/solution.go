package main

func shuffle(nums []int, n int) []int {
	var result []int
	for i, j := 0, n; j < len(nums); i, j = i+1, j+1 {
		result = append(result, nums[i])
		result = append(result, nums[j])
	}

	return result
}
