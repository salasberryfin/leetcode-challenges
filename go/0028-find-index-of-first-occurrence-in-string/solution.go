/*
Use two pointers: [start, start+len(needle)-1]
Iterate using the start pointer:
	- Check if start matches the first characters of needle
	- If not, move start to the right
	- If it matches:
		- if haystack[start:end-1] == needle -> return start
		- else -> move start to the right
*/
package main

func strStr(haystack string, needle string) int {
	result := -1

	n := len(needle)
	start := 0
	for start+n-1 < len(haystack) {
		end := start + n - 1
		if haystack[start] != needle[0] {
			start++
		} else {
			if haystack[start:end+1] == needle {
				return start
			} else {
				start++
			}
		}
	}

	return result
}
