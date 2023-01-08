/*
Input:
	points: array of points (x, y)
Result:
	Number of points that lie in the same line

Notes:
	- All points are unique
	- 1 <= len(points) <= 300

equation of line: y = mx + b

e.g. points = [[1,1], [2,2], [3,3]]

1. For each pair of points: get slope
2. map[slope]int = {
	slope: points in line
}
3. Create a new map with each iteration to avoid counting parallel lines (different line, same slope)
4. For every iteration, check if the number of points in the current slope is larger than the result value `max`

When calculating the slope:
	- Return inf for vertical lines
*/
package main

import (
	"math"
)

func getSlope(point1, point2 []int) float64 {
	if point1[0] != point2[0] {
		return float64(point2[1]-point1[1]) / float64(point2[0]-point1[0])
	} else {
		return math.Inf(1)
	}
}

func maxPoints(points [][]int) int {
	var max int

	for i := 0; i < len(points); i++ {
		var lines = map[float64]int{}
		for j := i + 1; j < len(points); j++ {
			slope := getSlope(points[i], points[j])
			lines[slope] += 1
			if lines[slope] > max {
				max = lines[slope]
			}
		}
	}

	return max + 1
}
