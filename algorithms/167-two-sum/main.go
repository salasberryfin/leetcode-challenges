package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		if numbers[l]+numbers[r] == target {
			return []int{
				l + 1,
				r + 1,
			}
		} else if numbers[l]+numbers[r] < target {
			l++
		} else {
			r--
		}
	}

	return []int{}
}

func twoSumInefficient(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target && i != j {
				return []int{
					i + 1,
					j + 1,
				}
			}
		}
	}

	return []int{}
}

func main() {
	inputs := [][]int{
		{2, 7, 11, 15},
		{2, 3, 4},
		{-1, 0},
	}

	targets := []int{
		9,
		6,
		-1,
	}

	var i int
	for i < len(inputs) {
		fmt.Println("Find in", inputs[i], "two nums that add up to", targets[i])
		//fmt.Println(twoSumInefficient(inputs[i], targets[i]))
		fmt.Println(twoSum(inputs[i], targets[i]))
		i++
	}
}
