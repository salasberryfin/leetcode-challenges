package main

import (
	"strings"
)

func reverseWords(s string) string {
	split := strings.Split(s, " ")

	var x int
	for x < len(split) {
		bytes := []byte(split[x])
		l, r := 0, len(bytes)-1
		for l < r {
			temp := bytes[l]
			bytes[l] = bytes[r]
			bytes[r] = temp
			l++
			r--
		}
		split[x] = string(bytes)
		x++
	}

	return strings.Join(split, " ")
}

func main() {
	inputs := []string{
		"Let's take LeetCode contest",
		"God Ding",
	}

	var x int
	for x < len(inputs) {
		reverseWords(inputs[x])
		x++
	}
}
