package main

import (
	"strconv"
)

func compress(chars []byte) int {
	r, l := 0, 0
	counter := 0
	ind := 0
	n := len(chars)
	for l < n && r < n {
		for chars[r] == chars[l] {
			counter++
			if r+1 >= len(chars) {
				break
			}
			r++
		}
		chars[ind] = chars[l]
		ind++
		new := []byte(strconv.Itoa(counter))
		if counter > 1 {
			for i := 0; i < len(new); i++ {
				chars[ind+i] = new[i]
			}
			ind += len(new)
		}
		l += counter
		r = l
		counter = 0
	}
	chars = chars[:ind]

	return len(chars)
}
