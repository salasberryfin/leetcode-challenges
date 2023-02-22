package main

type Interval struct {
	Start int
	End   int
}

func sort(left, right []int) []int {
	sorted := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if right[j] < left[i] {
			sorted = append(sorted, right[j])
			j++
		} else {
			sorted = append(sorted, left[i])
			i++
		}
	}

	for i < len(left) {
		sorted = append(sorted, left[i])
		i++
	}
	for j < len(right) {
		sorted = append(sorted, right[j])
		j++
	}

	return sorted
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	return sort(left, right)
}

func minMeetingRooms(intervals []*Interval) int {
	rooms := 0
	result := 0

	start, end := []int{}, []int{}
	for _, inter := range intervals {
		start = append(start, inter.Start)
		end = append(end, inter.End)
	}
	start = mergeSort(start)
	end = mergeSort(end)

	i, j := 0, 0
	for i < len(start) && j < len(end) {
		if start[i] < end[j] {
			rooms++
			i++
		} else {
			rooms--
			j++
		}
		if rooms > result {
			result = rooms
		}
	}

	return result
}
