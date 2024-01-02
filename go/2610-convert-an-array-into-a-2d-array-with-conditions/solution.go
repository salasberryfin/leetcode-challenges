package main

func findMatrix(nums []int) [][]int {
	result := [][]int{}

	numToApps := make(map[int]int)
	appsToNums := make(map[int][]int)
	for _, val := range nums {
		numToApps[val] += 1
		appsToNums[numToApps[val]] = append(appsToNums[numToApps[val]], val)
	}

	for _, i := range appsToNums {
		result = append(result, i)
	}

	return result
}
