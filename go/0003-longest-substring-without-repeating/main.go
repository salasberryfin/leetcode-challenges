package main

import "fmt"

func lengthOfLongestSubstrings(s string) int {
	occ := make(map[byte]int)
	var length int
	var maxLength int
	var start int
	var x int
	//for x := 0; x < len(s); x++ {
	for x < len(s) {
		//fmt.Println("Character:", string(s[x]))
		//fmt.Println("Current length is", length)
		char := s[x]
		if _, ok := occ[s[x]]; !ok {
			// this character has not been seen before
			length++
			if length > maxLength {
				//fmt.Println("Valid substring is now longer:", length)
				//fmt.Println("Found better solution")
				maxLength = length
			}
			occ[char]++
			x++
		} else {
			//fmt.Println("Found repetition")
			start++
			x = start
			length = 0
			occ = make(map[byte]int)
		}
		//occ[char]++
	}

	fmt.Println("Result :", maxLength)

	return maxLength
}

func main() {
	inputs := []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
		" ",
		"dvdf",
	}

	for _, x := range inputs {
		lengthOfLongestSubstrings(x)
	}
}
