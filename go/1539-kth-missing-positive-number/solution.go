package main

func findKthPositive(arr []int, k int) int {
	left, right := 0, len(arr)-1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		missing := arr[mid] - mid - 1
		if missing < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left + k
}
