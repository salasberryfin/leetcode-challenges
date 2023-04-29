package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	regexSubs := regexp.MustCompile(`[^a-z0-9]+`)
	s = regexSubs.ReplaceAllString(strings.ToLower(s), "")

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
