package main

func twoSum(nums []int, target int) []int {
	hashMap := map[int]int{}
	for i, num := range nums {
		if _, ok := hashMap[target-num]; ok {
			return []int{hashMap[target-num], i}
		}
		hashMap[num] = i
	}

	return []int{}
}
