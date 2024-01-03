/*
* Just need to keep track of the numbers of laser beams in each row and the previous non-empty row.
* `prev` is the number of lasers in the previous non-empty row.
* `prev` * `lasers` is the number of laser beams between both rows.
* Empty rows are just ignored.
 */
package main

func countLasers(str string) int {
	var lasers int
	for _, i := range str {
		if i == '1' {
			lasers++
		}
	}

	return lasers
}

func numberOfBeams(bank []string) int {
	var result int

	var prev int

	for _, i := range bank {
		lasers := countLasers(i)
		result += lasers * prev
		if lasers != 0 {
			prev = lasers
		}
	}

	return result
}
