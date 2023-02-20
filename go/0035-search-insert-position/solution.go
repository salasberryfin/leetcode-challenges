/*
Use a binary search method to find `target` in `nums`
Consider that `target` may not exist in `nums`
	- In this case, when right < left: return left
*/
package main

func binarySearch(nums []int, target, left, right int) int {
	if right < left {
		return left
	}
	mid := (right + left) / 2
	if nums[mid] == target {
		return mid
	} else if nums[mid] < target {
		return binarySearch(nums, target, mid+1, right)
	} else {
		return binarySearch(nums, target, left, mid-1)
	}
}

func searchInsert(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}
