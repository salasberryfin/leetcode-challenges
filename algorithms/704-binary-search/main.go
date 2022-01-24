package main

import "fmt"

func binarySearch(nums []int, target, l, h int) int {
	if h >= l {
		mid := (l + h) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			return binarySearch(nums, target, mid+1, h)
		} else if nums[mid] > target {
			return binarySearch(nums, target, l, mid-1)
		}
	}

	return -1
}

func search(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}

func main() {
	input := []int{-1, 0, 3, 5, 9, 12, 15}
	target := 2
	//target := 15
	fmt.Println("Result: ", search(input, target))
}
