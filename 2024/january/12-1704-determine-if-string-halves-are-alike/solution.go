package main

func countVowels(s string) int {
	var count int
	vowels := map[rune]bool{
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	for _, i := range s {
		if _, ok := vowels[i]; ok {
			count++
		}
	}

	return count
}

func halvesAreAlike(s string) bool {
	return countVowels(s[:len(s)/2]) == countVowels(s[len(s)/2:])
}
