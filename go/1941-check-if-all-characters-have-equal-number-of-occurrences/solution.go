package main

func areOccurrencesEqual(s string) bool {
	counter := map[byte]int{}
	for _, i := range s {
		counter[byte(i)]++
	}

	occ := counter[s[0]]
	for _, v := range counter {
		if v != occ {
			return false
		}
	}

	return true
}
