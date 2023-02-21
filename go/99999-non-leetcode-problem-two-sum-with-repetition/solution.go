package main

import "fmt"

func twoSumWithRepetition(nums []int, target int) [][]int {
	result := [][]int{}
	valueMap := map[int][]int{}

	for p, i := range nums {
		if _, ok := valueMap[target-i]; ok {
			result = append(result, []int{i, target - i})
		}
		valueMap[i] = append(valueMap[i], p)
	}

	fmt.Printf("Result %v\n", result)

	return result
}
