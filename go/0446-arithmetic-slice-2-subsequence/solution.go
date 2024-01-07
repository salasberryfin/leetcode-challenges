package main

func numberOfArithmeticSlices(nums []int) int {
	var result int
	dp := make([]map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = make(map[int]int)
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]
			if val, ok := dp[j][diff]; ok {
				dp[i][diff] += val + 1
				result += val
			} else {
				dp[i][diff] += 1
			}
		}
	}

	return result
}
