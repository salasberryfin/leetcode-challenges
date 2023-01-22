/*
Backtrack to find all possible substrings within the string and validate that each
of them is a palindrome

Keys:
	- The backtrack goes like:
		- for each starting position loop over all possible end positions, e.g. s[start:end]
		- check if each substring is a valid palindrome
		- if we get to the end of the string - len(s) - start <= 0 -> append to result
			- cannot append current directly because result would be changed whenever we change current
			- create a short-lived []string and copy the content of current
			- then append to result
*/
package main

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

func backtrack(s string, pos int, current []string, result *[][]string) {
	if len(s)-pos <= 0 {
		// need to create a new []string and copy content of current
		// if not, when changing current afterwars, result is changed as well
		newPalindrome := make([]string, len(current))
		copy(newPalindrome, current)
		*result = append(*result, newPalindrome)
	}

	for i := pos; i < len(s); i++ {
		if isPalindrome(s[pos : i+1]) {
			current = append(current, s[pos:i+1])
			backtrack(s, i+1, current, result)
			current = current[:len(current)-1]
		}
	}
}

func partition(s string) [][]string {
	var result [][]string

	backtrack(s, 0, []string{}, &result)

	return result
}
