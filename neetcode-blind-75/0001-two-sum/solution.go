package main

func twoSum(nums []int, target int) []int {
	values := map[int]int{}
	for i, v := range nums {
		if _, ok := values[target-v]; ok {
			return []int{values[target-v], i}
		}
		values[v] = i
	}

	return []int{}
}
