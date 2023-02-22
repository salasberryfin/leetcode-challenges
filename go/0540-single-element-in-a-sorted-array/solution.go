/*
Use two pointers: (m, m+1)
For each `m`, `m+1` must be equal
If it isn't -> return nums[m]
To consider the corner case of the last element being single:
	- return nums[len(nums)-1] if getting out of the loop without a result
*/
package main

func singleNonDuplicate(nums []int) int {
	for i := 0; i+1 < len(nums); i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}

	return nums[len(nums)-1]
}
