package main

func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		current := numbers[i] + numbers[j]
		if current == target {
			return []int{i + 1, j + 1}
		} else if target > current {
			i++
		} else {
			j--
		}
	}

	return []int{}
}
