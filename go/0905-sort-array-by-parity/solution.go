/*
First approach I took is doing merge sort and adapting to sorting on even/odd
This works but is slower than the simpler approach of just swapping elements in place
	- Use two pointers and, if left value is odd, swap with right value
	- Decrease counter of right value
	- If the left value is even, increse counter of left value
*/
package main

func deadSimple(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]%2 != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else {
			i++
		}
	}

	return nums
}

func sort(left, right []int) []int {
	i, j := 0, 0
	result := []int{}
	for i < len(left) && j < len(right) {
		if right[j]%2 != 0 && left[i]%2 != 0 {
			break
		}
		if left[i]%2 == 0 {
			result = append(result, left[i])
			i++
		}
		if right[j]%2 == 0 {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

func sortArrayByParity(nums []int) []int {
	return deadSimple(nums)
	//if len(nums) <= 1 {
	//	return nums
	//}

	//mid := len(nums) / 2
	//left := sortArrayByParity(nums[:mid])
	//right := sortArrayByParity(nums[mid:])

	//return sort(left, right)
}
