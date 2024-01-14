/*
* The key is counting the number of appearances of each character in each string.
* - Operation 1: reorder string.
*	Count occurrences of each character and compare the results for both words
* - Operation 2: reassign letters frequency.
*	Count the number of occurrences of each frequency, e.g. there two characters that appear 3 times,
*	one character that appears 2 times and one character that appears 1 time.
*	If the number of frequency occurrences is equal, we finally have to check if
*	the characters of both words are the same. If they aren't, not matter how many times
*	we swap characters, we won't be able to make the words equal.
*
 */
package main

import (
	"reflect"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	set1, set2 := make(map[byte]bool), make(map[byte]bool)
	countChars1, countChars2 := make(map[byte]int), make(map[byte]int)
	countFreq1, countFreq2 := make(map[int]int), make(map[int]int)
	for i := 0; i < len(word1); i++ {
		countChars1[word1[i]]++
		countChars2[word2[i]]++
		set1[word1[i]] = true
		set2[word2[i]] = true
	}
	if reflect.DeepEqual(countChars1, countChars2) {
		return true
	}
	for _, v := range countChars1 {
		countFreq1[v]++
	}
	for _, v := range countChars2 {
		countFreq2[v]++
	}
	return reflect.DeepEqual(countFreq1, countFreq2) && reflect.DeepEqual(set1, set2)
}
