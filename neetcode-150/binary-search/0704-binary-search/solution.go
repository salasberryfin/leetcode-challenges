package main

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (r + l) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}
