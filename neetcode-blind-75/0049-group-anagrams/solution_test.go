package main

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	strs := []string{
		"eat", "tea", "tan", "ate", "nat", "bat",
	}
	result := groupAnagrams(strs)
	fmt.Printf("The result is %v\n", result)
}
