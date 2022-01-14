package main

func merge(leftArray, rightArray []int) []int {
	size := len(leftArray) + len(rightArray)

	var arr []int
	arr = make([]int, size)

	// iterate over the given parts of the original array
	x, y, k := 0, 0, 0
	for x <= len(leftArray)-1 && y <= len(rightArray)-1 {
		if leftArray[x] <= rightArray[y] {
			arr[k] = leftArray[x]
			x++
		} else {
			arr[k] = rightArray[y]
			y++
		}
		k++
	}

	// check for remaining numbers in left array
	for x <= len(leftArray)-1 {
		arr[k] = leftArray[x]
		x++
		k++
	}

	// check for remaining numbers in right array
	for y <= len(rightArray)-1 {
		arr[k] = rightArray[y]
		y++
		k++
	}

	//fmt.Println("Result:", arr)

	return arr
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2

	// split array
	left, right := nums[:mid], nums[mid:]

	// recursively split array until len(array)=1
	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left, right)
}
