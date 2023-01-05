/*
Use a map to store the count of each unique element: map[value]count
Iterate over each key of the map.
	1. If the number of appearances of the character is odd: odd++
	2. We can only have one character that appears odd times: in the middle.
	3. Then, if odd != 0 -> result = result - odd + 1
*/
package main

import (
	"strings"
)

func longestPalindrome(s string) int {
	result := len(s)
	var odd int
	uniqueValues := map[rune]int{}
	for _, i := range s {
		occurrences := strings.Count(s, string(i))
		uniqueValues[i] = occurrences
	}

	for _, v := range uniqueValues {
		if v%2 != 0 {
			odd++
		}
	}

	var subtract int
	if odd != 0 {
		subtract = odd - 1
	}

	return result - subtract
}
