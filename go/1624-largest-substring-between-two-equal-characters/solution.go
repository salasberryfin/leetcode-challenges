/*
* Keep track of the positions for each characters in s
* Find the largest distances between positions
* The largest distances defaults to -1
 */
package main

func maxLengthBetweenEqualCharacters(s string) int {
	longest := -1
	positions := make(map[rune][]int)
	for pos, val := range s {
		positions[val] = append(positions[val], pos)
		if positions[val][len(positions[val])-1]-positions[val][0]-1 > longest {
			longest = positions[val][len(positions[val])-1] - positions[val][0] - 1
		}
	}

	return longest
}
