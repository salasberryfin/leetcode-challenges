/*
Use binary search to navigate through all possible solutions

- The key idea is that the number k -per-hour eating speed- must be an integer
between 1 and the largest number in piles
- Binary Search between those two elements:
  - Check if the current `mid` is valid: `hours <= h`
  - if so -> move left
  - else -> move right

- Return left
*/
package main

func check(piles []int, mid, h int) bool {
	hours := 0
	for _, i := range piles {
		hours += (i + (mid - 1)) / mid
	}
	if hours <= h {
		return true
	}

	return false
}

func max(piles []int) int {
	max := 0
	for _, i := range piles {
		if i > max {
			max = i
		}
	}

	return max
}

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, max(piles)
	for left < right {
		mid := (left + right) / 2
		if check(piles, mid, h) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}
