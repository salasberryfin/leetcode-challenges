/*
Use Binary Search to go over all `moments`: between 0 and the maximum time
For each time moment:
	- count the number of trips until then
	- if larger or equal to `totalTrips` -> move right to the left
	- else -> move left to the right
Return left
*/
package main

import (
	"math"
)

func find(time []int, moment, target int) bool {
	trips := 0
	for _, i := range time {
		trips += moment / i
		if trips >= target {
			return true
		}
	}

	return false
}

func minimumTime(time []int, totalTrips int) int64 {
	left, right := 0, math.MaxInt

	for left <= right {
		mid := (left + right) / 2
		if find(time, mid, totalTrips) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return int64(left)
}
