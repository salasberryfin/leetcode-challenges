package main

import "fmt"

func isPalindrome(x int) bool {
	strX := fmt.Sprint(x)

	i, j := 0, len(strX)-1
	for i < j {
		if strX[i] != strX[j] {
			return false
		}
		i++
		j--
	}

	return true
}
