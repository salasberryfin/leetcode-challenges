package main

import "fmt"

func sort(left, right []byte) []byte {
	result := make([]byte, len(left)+len(right))

	var x, y, k int
	for x <= len(left)-1 && y <= len(right)-1 {
		if right[y] < left[x] {
			result[k] = right[y]
			y++
		} else {
			result[k] = left[x]
			x++
		}
		k++
	}

	for x <= len(left)-1 {
		result[k] = left[x]
		x++
		k++
	}

	for y <= len(right)-1 {
		result[k] = right[y]
		y++
		k++
	}

	return result
}

func mergeSort(arr []byte) []byte {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return sort(left, right)
}

func findTheDifference(s string, t string) byte {
	var result byte

	sortedS := mergeSort([]byte(s))
	sortedT := mergeSort([]byte(t))

	var x int
	for x <= len(sortedT) {
		if x >= len(sortedS) {
			return sortedT[x]
		}
		if sortedS[x] != sortedT[x] {
			return sortedT[x]
		}
		x++
	}

	fmt.Println("Sorted S bytes:", sortedS)
	fmt.Println("T bytes:", sortedT)

	return result
}

func main() {
	//s := "abcd"
	//t := "abcde"
	s := ""
	t := "y"

	fmt.Println(findTheDifference(s, t))
}
