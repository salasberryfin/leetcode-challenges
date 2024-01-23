package main

func findUniqueChars(s string) int {
	var charArr [26]int
	for _, char := range s {
		charArr[char-'a']++
		if charArr[char-'a'] > 1 {
			return 0
		}
	}
	var count int
	for _, val := range charArr {
		if val == 1 {
			count++
		}
	}

	return count
}

func dfs(pos int, s string, arr []string) int {
	if pos == len(arr) {
		return findUniqueChars(s)
	}

	valid := dfs(pos+1, s+arr[pos], arr)
	nonValid := dfs(pos+1, s, arr)

	if valid > nonValid {
		return valid
	} else {
		return nonValid
	}
}

func maxLength(arr []string) int {
	return dfs(0, "", arr)
}
