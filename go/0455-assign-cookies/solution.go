/*
* Sort both slices: cookies and greed factor
* Use two pointers to check which cookie satisfies each kid
 */
package main

func sort(l, r []int) []int {
	result := make([]int, len(l)+len(r))
	i, j, k := 0, 0, 0
	for i < len(l) && j < len(r) {
		if r[j] < l[i] {
			result[k] = r[j]
			j++
		} else {
			result[k] = l[i]
			i++
		}
		k++
	}

	for i < len(l) {
		result[k] = l[i]
		i++
		k++
	}
	for j < len(r) {
		result[k] = r[j]
		j++
		k++
	}

	return result
}

func mergeSort(m []int) []int {
	if len(m) <= 1 {
		return m
	}

	mid := len(m) / 2
	left := mergeSort(m[:mid])
	right := mergeSort(m[mid:])

	return sort(left, right)
}

func findContentChildren(g []int, s []int) int {
	sortedG := mergeSort(g)
	sortedS := mergeSort(s)

	var satisfiedKids int
	i, j := 0, 0
	for i < len(sortedG) && j < len(sortedS) {
		if sortedG[i] <= sortedS[j] {
			satisfiedKids++
			i++
		}
		j++
	}

	return satisfiedKids
}
