package main

import "fmt"

func binarySearch(nums []int, target, l, h int) int {
	mid := (h + l) / 2
	if h > l {
		if nums[mid] == target {
			// found position
			return mid
		} else if nums[mid] > target {
			// target should be in the lower positions
			return binarySearch(nums, target, l, mid-1)
		} else if nums[mid] < target {
			// target should be in the higher positions
			return binarySearch(nums, target, mid+1, h)
		}
	}

	if nums[mid] < target {
		return mid + 1
	}

	return mid
}

func searchInsert(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}

func main() {
	input := []int{1, 3, 5, 6}
	//input := []int{3, 5, 7, 9, 10}
	target := 0
	//target := 0
	result := searchInsert(input, target)
	fmt.Println("Result ", result)
}
