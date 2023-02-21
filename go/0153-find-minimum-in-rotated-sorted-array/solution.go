/*
Keys:
	- first check if the array is not rotated or if it has been rotated len(nums) times
		- in that case, nums[0] < nums[len(nums)-1] -> is ordered: return nums[0]
	- then we can concentrate on rotated arrays only
	- use binary search to navigate the elements
	- the key idea is that rotated arrays have two parts ordered independently and a pivot point
	- check the middle point in the binary search against the first item
		- if nums[0] <= nums[mid] -> this is the larger side of the array, move right
		- if nums[0] > nums[mid] -> this is the smaller side of the array, move left
	- keep a `min` variable and update on each cycle of the binary search if
		- nums[mid] < min -> min = nums[mid]
*/
package main

func findMin(nums []int) int {
	// non-rotated array
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	min := nums[0]
	left, right := 0, len(nums)-1
	// loop binary search
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < min {
			min = nums[mid]
		}

		if nums[0] <= nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return min
}
