package main

import (
	"fmt"
	"math"
)

func bubbleSort(nums []int) []int {
	// inefficient
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				// switch positions i and i+1
				curr := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = curr
			}
		}
	}

	return nums
}

func sortedSquares(nums []int) []int {
	var squares []int
	for _, x := range nums {
		power := int(math.Pow(float64(x), float64(2)))
		squares = append(squares, power)
	}

	var sorted []int
	sorted = bubbleSort(squares)
	//sorted = mergeSort(squares)

	return sorted
}

func main() {
	inputs := [][]int{
		{-4, -1, 0, 3, 10},
		{-7, -3, 2, 3, 11},
		{3, 1, 2},
	}
	for _, input := range inputs {
		fmt.Println("Result for input", input, "is", sortedSquares(input))
	}
}
