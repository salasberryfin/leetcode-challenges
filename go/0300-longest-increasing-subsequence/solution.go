/*
Hard problem to understand at first sight

The key is to iterate backwards and keep track of each of the longest increasing subsequence (LIS) for each element
- e.g. the last element will always have a LIS=1
- the second last will be 1+LIS[last] if last element > second last element
- and keep going...

Start by creating a LIS array to store each element's LIS, and initiliaze to 1s
When checking an element that has multiple other elements to the right side of the array:
- there are multiple options to choose from
- check all of them and keep track of the max LIS of all of them
*/
package main

func lengthOfLIS(nums []int) int {
	var result int
	lisArray := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		lisArray[i] = 1
	}

	// iterate backwards
	for i := len(nums) - 1; i >= 0; i-- {
		current := nums[i]
		// check each possibility
		best := 0
		for j := i + 1; j < len(lisArray); j++ {
			candidate := nums[j]
			if current < candidate {
				if lisArray[j] > best {
					best = lisArray[j]
				}
			}
		}
		lisArray[i] += best
		if lisArray[i] > result {
			result = lisArray[i]
		}
	}

	return result
}
