package main

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sOccurrences, tOccurrences := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(s); i++ {
		sOccurrences[s[i]]++
		tOccurrences[t[i]]++
	}
	for i := range sOccurrences {
		if sOccurrences[i] != tOccurrences[i] {
			return false
		}
	}

	return true
}
