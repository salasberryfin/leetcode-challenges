/*
Create a hashmap with all the unique values needed to generate all numbers
That is, all the unique letter plus the special characters with prefix [I, X, C]
	1. Iterate over each of the integers
	2. If num/integer is more than 0: we add as many characters of the roman value as needed
	3. Subtract the value from num
	4. Keep iterating until num == 0
*/
package main

func intToRoman(num int) string {
	var result string

	romans := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	if _, ok := romans[num]; ok {
		return romans[num]
	}

	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for _, val := range values {
		if num == 0 {
			break
		}
		if num/val != 0 {
			for i := 0; i < num/val; i++ {
				result += romans[val]
			}
			num -= (num / val) * val
		}
	}

	return result
}
