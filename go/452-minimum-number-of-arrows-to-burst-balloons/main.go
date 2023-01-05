/*
Balloons hell:
I was completely lost in this one.
Hints:
	1. Order the array of positions: I use mergeSort and adapt for 2d arrays
	2. Go over each element and check if it is contained in the next one
		This means there's only one arrow needed for all the contained balloons
	3. If it is not contained, try the second element of the current balloons
	4. If it still is not contained, find the minimun value between previous[1] and current[1]
*/

package main

func sort(left, right [][]int) [][]int {
	result := make([][]int, len(left)+len(right))

	var i, j, k int
	for i < len(left) && j < len(right) {
		if left[i][0] < right[j][0] {
			result[k] = left[i]
			i++
		} else {
			result[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		result[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		result[k] = right[j]
		j++
		k++
	}

	return result
}

func mergeSort(slice [][]int) [][]int {
	lenSlice := len(slice)
	if lenSlice == 1 {
		return slice
	}

	middle := lenSlice / 2
	left := mergeSort(slice[:middle])
	right := mergeSort(slice[middle:])

	slice = sort(left, right)

	return slice
}

func findMinArrowShots(points [][]int) int {
	var arrows int
	balloons := len(points)
	if balloons == 1 {
		return 1
	}

	sorted := mergeSort(points)

	arrows = 1

	i := 1
	var shot []int
	shot = sorted[0]

	for i < len(sorted) {
		if sorted[i][1] <= shot[1] && sorted[i][0] >= shot[0] {
			//fmt.Printf("Have to find min(%v, %v): %d\n", shot, sorted[i], sorted[i][1])
			shot = sorted[i]
		} else if shot[1] < sorted[i][0] || shot[1] > sorted[i][1] {
			//fmt.Printf("%d is not contained in %d, must look for a different shot %d\n", shot, sorted[i], sorted[i])
			shot = sorted[i]
			arrows++
		}
		i++
	}

	return arrows
}
