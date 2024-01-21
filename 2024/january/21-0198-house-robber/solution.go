package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	house1, house2 := 0, 0
	for i := 0; i < len(nums); i++ {
		current := max(house1+nums[i], house2)
		house1 = house2
		house2 = current
	}

	return house2
}
