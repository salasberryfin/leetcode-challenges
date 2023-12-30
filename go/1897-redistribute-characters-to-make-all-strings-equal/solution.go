/*
* This can be reduced to a counting characters problem.
* Knowing we can do any number of permutations in the input words, we only
* need to count the number of appearances of each character.
* - This number, must either be equal to the number of words or divisible (%) by
* the number of items in `words`, so that characters can be evenly distributed.
 */
package main

func countAppearances(words []string) map[rune]int {
	m := make(map[rune]int)
	for _, word := range words {
		for _, char := range word {
			m[char] += 1
		}
	}

	return m
}

func makeEqual(words []string) bool {
	count := countAppearances(words)
	for _, i := range count {
		if i%len(words) != 0 {
			return false
		}
	}

	return true
}
