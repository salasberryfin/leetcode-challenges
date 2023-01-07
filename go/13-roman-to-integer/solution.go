/*
Create a hashmap with all the roman numbers
Start from the leftmost element of the string and go right
	1. Get value of current character from map
	2. If current value is greater than previous value: we subtract
	3. Else: we add
	4. Update previous value
*/
package main

func romanToInt(s string) int {
	romans := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	if _, ok := romans[s]; ok {
		return romans[s]
	}

	var number int
	var previous int
	for i := len(s) - 1; i >= 0; i-- {
		current := romans[string(s[i])]
		if previous > current {
			number -= current
		} else {
			number += current
		}
		previous = current
	}

	return number
}
