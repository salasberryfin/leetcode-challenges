package main

func mergeSort(l, r []int) []int {
	sorted := make([]int, len(l)+len(r))

	var i, j, k int
	for i < len(l) && j < len(r) {
		if r[j] < l[i] {
			sorted[k] = r[j]
			j++
		} else {
			sorted[k] = l[i]
			i++
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

func sort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := sort(nums[:mid])
	right := sort(nums[mid:])

	return mergeSort(left, right)
}

func buyChoco(prices []int, money int) int {
	ordered := sort(prices)
	currentMoney := money
	var chocolates int
	for _, i := range ordered {
		currentMoney -= i
		if currentMoney < 0 {
			return money
		}
		chocolates += 1
		if chocolates == 2 {
			return currentMoney
		}
	}

	return 0
}
