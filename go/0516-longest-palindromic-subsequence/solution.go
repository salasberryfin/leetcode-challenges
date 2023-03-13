/*
Follow a dynamic programming approach - recursively
The key idea is that we can find a repetition pattern to get results
If opposite items of a substring match: we add 2 and check the substring to the left and right
If opposite items are different: we either want to go to the left or to the right, so we take the max()
Create a matrix of longest palindromic subsequence for each combination of (left, right)

	For each left, right
	if s[left] == s[right]:
		lps = 2+l[left+1][right-1]
	else:
		max(l[left+1][right], l[left][right-1])

It is necessary to avoid duplicating results so we check if the (left, right) combination is already
in the results matrix
*/
package main

func max(l, r int) int {
	if l > r {
		return l
	}

	return r
}

func dynamic(s string, values [][]int, left, right int) int {
	if left > right {
		return 0
	} else if left == right {
		return 1
	}

	// check that this combination of (left, right) has not been checked yet
	if values[left][right] == 0 {
		if s[left] == s[right] {
			values[left][right] = 2 + dynamic(s, values, left+1, right-1)
		} else {
			values[left][right] = max(dynamic(s, values, left+1, right), dynamic(s, values, left, right-1))
		}
	}

	return values[left][right]
}

func longestPalindromeSubseq(s string) int {
	var result int
	n := len(s)
	values := make([][]int, n)
	for i := range values {
		values[i] = make([]int, n)
	}
	result = dynamic(s, values, 0, len(s)-1)

	return result
}
