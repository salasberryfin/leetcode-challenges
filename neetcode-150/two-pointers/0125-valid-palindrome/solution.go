package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	nonAlphanumeric := regexp.MustCompile("[^a-zA-Z0-9]+")
	s = strings.ToLower(nonAlphanumeric.ReplaceAllString(s, ""))

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}
