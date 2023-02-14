package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash := map[rune]int{}
	for _, i := range s {
		hash[i] = hash[i] + 1
	}
	for _, i := range t {
		if val, ok := hash[i]; !ok || val == 0 {
			return false
		}
		hash[i] = hash[i] - 1
	}

	return true
}
