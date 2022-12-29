package main

import "fmt"

func isPalindrome(x int) bool {
	strX := fmt.Sprint(x)
	i, j := 0, len(strX)
	for i < j {
		if strX[i] != strX[j-1] {
			return false
		}
		i++
		j--
	}

	return true
}
