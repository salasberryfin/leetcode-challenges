/*
Use backtrack to generate only valid solutions:
	1. Keep track of open parenthesis (left) and close parenthesis (close)
		a. This allows to keep track of what's valid
	2. If we still have parenthesis left (<n) open a new one
	3. If we have opened parenthesis but not closing, close
	4. When reaching the end (2*n) join all characters of the slice into a string
	and return
*/
package main

import (
	"fmt"
	"strings"
)

func backtrack(s []string, left, right, n int) []string {
	var result []string
	if len(s) == 2*n {
		result = append(result, strings.Join(s, " "))
		return result
	}
	if left < n {
		s = append(s, "(")
		result = append(result, backtrack(s, left+1, right, n)...)
		s = s[:len(s)-1]
	}
	if right < left {
		s = append(s, ")")
		result = append(result, backtrack(s, left, right+1, n)...)
		s = s[:len(s)-1]
	}

	return result
}

func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{
			"()",
		}
	}
	solution := backtrack([]string{}, 0, 0, n)
	fmt.Println(solution)

	var result []string
	for _, i := range solution {
		current := strings.Replace(i, " ", "", -1)
		result = append(result, current)
	}

	return result
}
