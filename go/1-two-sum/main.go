package main

import (
	"fmt"
	"time"
)

type item struct {
	pos int
	val int
}

func sort(left, right []item) []item {
	result := make([]item, len(left)+len(right))

	var x, y, k int
	for x <= len(left)-1 && y <= len(right)-1 {
		if right[y].val < left[x].val {
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

func mergeSort(arr []item) []item {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return sort(left, right)
}

func contains(value int, slice []int) (int, bool) {
	for i, v := range slice {
		if v == value {
			return i, true
		}
	}

	return 0, false
}

func twoSum(nums []int, target int) []int {
	var itemSlice []item
	// keep track of the index of each item
	for i := 0; i < len(nums); i++ {
		itemSlice = append(itemSlice, item{pos: i, val: nums[i]})
	}
	// first sort the array
	sorted := mergeSort(itemSlice)
	i, j := 0, len(sorted)-1
	for i < j {
		candidate := sorted[i].val + sorted[j].val
		if candidate == target {
			return []int{sorted[i].pos, sorted[j].pos}
		} else if candidate > target {
			j--
		} else {
			i++
		}
	}

	return []int{}
}

// naive solution 2
func naive2TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		ind, cont := contains(target-nums[i], nums[i+1:])
		offset := ind + i + 1
		if cont {
			return []int{i, offset}
		}
	}

	return []int{}
}

// naive solution
func naiveTwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for k := i + 1; k < len(nums); k++ {
			if nums[i]+nums[k] == target {
				return []int{i, k}
			}
		}
	}

	return []int{}
}

func main() {
	example1 := []int{
		2, 7, 11, 15,
	}
	example2 := []int{
		3, 2, 4,
	}
	example3 := []int{
		3, 3,
	}
	examples := [][]int{
		example1,
		example2,
		example3,
	}
	targets := []int{
		9,
		6,
		6,
	}
	for i := 0; i < len(examples); i++ {
		// start by ordering the slice
		start := time.Now()
		resultNaive1 := naiveTwoSum(examples[i], targets[i])
		fmt.Printf("Naive solution 1 took %v\n", time.Since(start))
		fmt.Printf("Result [naive solution 1] of adding two elements of %v to get %d: %v\n", examples[i], targets[i], resultNaive1)
		start = time.Now()
		resultNaive2 := naive2TwoSum(examples[i], targets[i])
		fmt.Printf("Naive solution 2 took %v\n", time.Since(start))
		fmt.Printf("Result of adding two elements of %v to get %d: %v\n", examples[i], targets[i], resultNaive2)
		start = time.Now()
		result := twoSum(examples[i], targets[i])
		fmt.Printf("Optimal solution took %v\n", time.Since(start))
		fmt.Printf("Result of adding two elements of %v to get %d: %v\n", examples[i], targets[i], result)
	}

}
