package main

import (
	"fmt"
	"reflect"
)

// use a hash table (Go map) to store occurrences of each character in
// string and substring
func hash(s string) map[byte]int {
	a := make(map[byte]int)

	var i int
	for i < len(s) {
		a[s[i]]++
		i++
	}

	return a
}

func checkInclusion(s1 string, s2 string) bool {
	a1 := hash(s1)

	for x := 0; x <= len(s2)-len(s1); x++ {
		subs := s2[x : len(s1)+x]
		subHash := hash(subs)
		if reflect.DeepEqual(a1, subHash) {
			return true
		}
	}

	return false
}

func main() {
	inputs := [][]string{
		{"ab", "eidbaooo"},
		{"ab", "eidboaoo"},
		{"adc", "dcda"},
	}

	for _, x := range inputs {
		fmt.Println(checkInclusion(x[0], x[1]))
	}
}
