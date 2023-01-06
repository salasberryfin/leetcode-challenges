/*
1833 Maximum ice cream bars
An easy one.
	1. Order the slice of costs.
	2. Iterate until no more coins can be spent to buy next ice cream.
*/
package main

func mergeSort(l, r []int) []int {
	sorted := make([]int, len(l)+len(r))

	var i, j, k int
	for i < len(l) && j < len(r) {
		if l[i] < r[j] {
			sorted[k] = l[i]
			i++
		} else {
			sorted[k] = r[j]
			j++
		}
		k++
	}

	for i < len(l) {
		sorted[k] = l[i]
		i++
		k++
	}
	for j < len(r) {
		sorted[k] = r[j]
		j++
		k++
	}

	return sorted
}

func sort(slice []int) []int {
	if len(slice) == 1 {
		return slice
	}

	left := sort(slice[0 : len(slice)/2])
	right := sort(slice[len(slice)/2:])

	return mergeSort(left, right)
}

func maxIceCream(costs []int, coins int) int {
	var maximum int

	sortedCosts := sort(costs)

	if sortedCosts[0] > coins {
		return 0
	}

	for _, cost := range sortedCosts {
		if coins-cost < 0 {
			return maximum
		}
		coins -= cost
		maximum++
	}

	return maximum
}
