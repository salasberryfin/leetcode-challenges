/*
* The idea is to keep track of the number of zeros from the left and the number of ones from the right.
*	- Since both subarrays must have at least one element, we start from the second element from the left and the second element from the right.
*	- There is a special case when the length of the string is 2, in which case we just check if the first element is 0 and the second element is 1 because there are no other subarays.
 */
package main

func maxScore(s string) int {
	max := 0
	if len(s) == 2 {
		if s[0] == '0' {
			max++
		}
		if s[1] == '1' {
			max++
		}
		return max
	}
	zeros, ones := make([]int, len(s)), make([]int, len(s))

	zerosPrev, onesPrev := 0, 0
	for i, j := 0, len(s)-1; i < len(s) && j >= 0; i, j = i+1, j-1 {
		if s[i] == '0' {
			zeros[i] += zerosPrev + 1
		} else {
			zeros[i] = zerosPrev
		}
		if s[j] == '1' {
			ones[j] += onesPrev + 1
		} else {
			ones[j] = onesPrev
		}
		zerosPrev, onesPrev = zeros[i], ones[j]
	}
	for x := 1; x < len(zeros)-1; x++ {
		if zeros[x]+ones[x] > max {
			max = zeros[x] + ones[x]
		}
	}

	return max
}
