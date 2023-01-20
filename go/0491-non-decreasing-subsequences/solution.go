/*
Need to review Backtracking.
Couldn't come up with anything close to the solution until I looked at the
official solution
*/
package main

func backtrack(nums []int, sl []int, l int, res *[][]int) {
	if len(sl) > 1 {
		cp := make([]int, len(sl))
		copy(cp, sl)
		*res = append(*res, cp)
	}
	seen := make(map[int]bool, len(nums)-l)
	for r := l; r < len(nums); r++ {
		if (l > 0 && nums[r] < nums[l-1]) || seen[nums[r]] {
			continue
		}
		seen[nums[r]] = true
		backtrack(nums, append(sl, nums[r]), r+1, res)
	}
}

func findSubsequences(nums []int) [][]int {
	result := [][]int{}
	backtrack(nums, []int{}, 0, &result)

	return result
}
