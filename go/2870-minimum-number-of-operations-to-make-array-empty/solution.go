/*
If the number of occurrences is not divisible by 3, we just need to
add an extra operation for 2: round up v/3.
If it is divisible by 3, we just need the result of v/3.
*/
package main

import (
	"math"
)

func minOperations(nums []int) int {
	var operations int

	occurrences := make(map[int]int)
	for _, i := range nums {
		occurrences[i]++
	}
	for _, v := range occurrences {
		if v < 2 {
			return -1
		}

		if v%3 != 0 {
			operations += int(math.Ceil(float64(v) / 3))
		} else {
			operations += v / 3
		}
	}

	return operations
}
