package main

import (
	"fmt"
	"math"
)

func rotate(nums []int, k int) {
	fmt.Println("Rotating", nums, "with a right shift of", k)
	if len(nums) <= 1 {
		return
	}
	result := append([]int{}, nums...)
	for i, x := range result {
		if i+k < len(nums) {
			nums[i+k] = x
		} else {
			nums[int(math.Abs(float64((i+k)%(len(nums)))))] = x
		}
	}

	//fmt.Printf("Result address: %p\n", &result)
	//fmt.Printf("Original Nums address: %p\n", &nums)
	//nums = result
	//fmt.Printf("Post Nums address: %p\n", &nums)
	fmt.Println("result:", result)
	fmt.Println()
}

func main() {
	inputs := [][]int{
		{1, 2, 3, 4, 5, 6, 7},
		{1, -100, 3, 99},
		{1, 2},
	}
	ks := []int{
		3,
		2,
		1,
	}
	for _, input := range inputs {
		for _, k := range ks {
			rotate(input, k)
		}
	}
}
