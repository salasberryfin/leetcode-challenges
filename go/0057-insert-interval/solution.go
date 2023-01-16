/*
This is harder than it looks
The idea is to insert the new interval in the intervals slice while keeping the
order based on the interval start
Two different ways of solving this sorting problem are:
	- Linear search: iterate over the elements in intervals and find the first element
	with a start larger than the start of the new interval. Insert new interval
	- Binary search: the process is almost identical but we find the position
	where the new interval is inserted by doing a Binary search, which improves performance
	for larger inputs
*/
package main

func binarySearch(intervals [][]int, val []int) int {
	// this binary search allows to get the first start in intervals that
	// is larger than the start of newInterval, which is the position where
	// the new interval will be inserted in the slice
	left, right := 0, len(intervals)

	for left != right {
		mid := (left + right) / 2
		if val[0] >= intervals[mid][0] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return right
}

func linearSearch(intervals [][]int, newIntervals []int) int {
	for i, v := range intervals {
		if v[0] >= newIntervals[0] {
			return i
		}
	}

	return len(intervals)
}

func insert(intervals [][]int, newInterval []int) [][]int {
	intervalsLen := len(intervals)

	if intervalsLen == 0 {
		return [][]int{newInterval}
	}

	updatedIntervals := [][]int{}
	position := binarySearch(intervals, newInterval)
	updatedIntervals = append(updatedIntervals, intervals[:position]...)
	updatedIntervals = append(updatedIntervals, newInterval)
	updatedIntervals = append(updatedIntervals, intervals[position:]...)

	for i := 0; i < len(updatedIntervals); i++ {
		interval := updatedIntervals[i]
		for j := i + 1; j < len(updatedIntervals); j++ {
			if interval[1] >= updatedIntervals[j][0] {
				end := interval[1]
				if interval[1] < updatedIntervals[j][1] {
					end = updatedIntervals[j][1]
				}
				updatedIntervals = append(updatedIntervals[:i], updatedIntervals[j:]...)
				updatedIntervals[i] = []int{interval[0], end}
				i--
				break
			}
		}
	}

	return updatedIntervals
}
