package main

func twoSum(nums []int, target int) []int {
	// store values and index in hash map
	// iterate once over the nums array and check if target-nums[i] exists in the hash map -> this is the result

	hashMap := make(map[int]int)
	for i, num := range nums {
		if _, ok := hashMap[target-num]; ok {
			return []int{hashMap[target-num], i}
		}
		hashMap[num] = i
	}

	return []int{}
}
