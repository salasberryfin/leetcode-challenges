package main

func minSteps(s string, t string) int {
	var result int

	count := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		count[s[i]]++
		count[t[i]]--
	}

	for _, v := range count {
		if v > 0 {
			result += v
		}
	}

	return result
}
