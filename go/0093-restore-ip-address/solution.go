/*
Backtracking to validate each solution

Keys:
	- Only 4 blocks between 0 and 255 (both included) are valid.
	- Blocks that start with 0 and are not 0 (len(block) > 1) are invalid.
	- Only three dots can be used.

Backtrack through all possible elements and validate if the candidate IP is valid
using the isIp function
*/
package main

import (
	"strconv"
)

func isIp(candidate string) bool {
	if len(candidate) > 3 {
		return false
	}
	if candidate[0] == '0' && len(candidate) > 1 {
		return false
	}
	candNum, _ := strconv.Atoi(candidate)

	return candNum <= 255 && candNum >= 0
}

func backtrack(s, candidate string, start, dots int, result *[]string) {
	remaining := len(s) - start
	if remaining <= 0 {
		return
	}
	if dots == 0 {
		if isIp(s[start:]) {
			*result = append(*result, candidate+s[start:])
		}
	}
	for i := start + 1; i < len(s); i++ {
		if isIp(s[start:i]) {
			backtrack(s, candidate+s[start:i]+".", i, dots-1, result)
		}
	}
}

func restoreIpAddresses(s string) []string {
	var result []string

	if len(s) > 12 {
		return []string{}
	}
	backtrack(s, "", 0, 3, &result)

	return result
}
