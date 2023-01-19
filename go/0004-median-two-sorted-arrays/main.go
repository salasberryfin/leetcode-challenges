package main

import "fmt"

func mergeSort(left, right []int) []int {
	x, y := 0, 0
	var result []int
	for x < len(left) && y < len(right) {
		if left[x] <= right[y] {
			result = append(result, left[x])
			x++
		} else if right[y] < left[x] {
			result = append(result, right[y])
			y++
		}
	}

	for x < len(left) {
		result = append(result, left[x])
		x++
	}
	for y < len(right) {
		result = append(result, right[y])
		y++
	}

	return result
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var result float64
	merged := mergeSort(nums1, nums2)
	fmt.Printf("Merged array: %v\n", merged)

	if len(merged)%2 != 0 {
		// if odd: middle element
		result = float64(merged[len(merged)/2])
	} else {
		// if even: sum two middle elements
		sum := merged[len(merged)/2] + merged[len(merged)/2-1]
		result = float64(float64(sum) / 2)
	}
	fmt.Printf("Median: %f\n", result)

	return result
}

func main() {
	inputs := [][][]int{
		{
			{1, 3},
			{2},
		},
		{
			{1, 2},
			{3, 4},
		},
	}

	for k, i := range inputs {
		fmt.Printf("Input %d: %v:\n", k, i)
		findMedianSortedArrays(i[0], i[1])
	}
}
