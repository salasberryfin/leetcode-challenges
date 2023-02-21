/*
The idea behind this solution is that binary representation of numbers follows
a pattern
Each number has as many 1s as: 1 + number of 1s of the number in position current-most significant bit

This means:
	- ones[1] = 1 + ones[1 - 1]
	- ones[2] = 1 + ones[2 - 1]
	- ones[3] = 1 + ones[3 - 2]
	- ones[4] = 1 + ones[4 - 4]
	- ones[5] = 1 + ones[5 - 4]
	- ones[6] = 1 + ones[6 - 4]
	...
	- ones[8] = 1 + ones[8 - 8]
	- ones[9] = 1 + ones[9 - 8]

With significant bits being the powers of 2: [1, 2, 4, 6, 8, 16,...]

Iterating over all elements until n:
	- Check if the value of the most significant bit needs update
	- Calculate the number of ones
*/
package main

func countBits(n int) []int {
	result := make([]int, n+1)
	result[0] = 0
	significant := 1
	// # of 1s = 1 + number of 1s[current-significant]
	for i := 1; i <= n; i++ {
		if i == significant*2 {
			significant = i
		}
		result[i] = 1 + result[i-significant]
	}

	return result
}
