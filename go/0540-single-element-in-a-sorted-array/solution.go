/*
The idea is dead simple if you think about it: all elements must be in pairs
	- postions even|odd must be the same
	- when a new element in an even position does not match the element in the following odd position
		- return this element
	- if we reach the end and haven't found any element, that means the last element is the single one
		- this applies also when len(nums) == 1
*/
package main

func singleNonDuplicate(nums []int) int {
	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}

	return nums[len(nums)-1]
}
