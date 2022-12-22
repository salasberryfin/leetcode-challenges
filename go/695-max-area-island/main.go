package main

import (
	"fmt"
)

func neighbors(row, col, m, n int, grid [][]int) [][]int {
	var result [][]int

	// up
	if row+1 < m {
		if grid[row+1][col] == 1 {
			result = append(result, []int{row + 1, col})
		}
	}
	// down
	if row-1 >= 0 {
		if grid[row-1][col] == 1 {
			result = append(result, []int{row - 1, col})
		}
	}
	// right
	if col+1 < n {
		if grid[row][col+1] == 1 {
			result = append(result, []int{row, col + 1})
		}
	}
	// left
	if col-1 >= 0 {
		if grid[row][col-1] == 1 {
			result = append(result, []int{row, col - 1})
		}
	}

	return result
}

func breadthSearch(row, col, m, n int, grid [][]int) int {
	var counter int

	queue := make([][]int, 0)
	queue = append(queue, []int{row, col})
	visited := map[int]bool{}
	visited[queue[0][0]+queue[0][1]*m] = true

	for len(queue) > 0 {
		counter++
		node := queue[0]
		//fmt.Println("Node", node)
		queue = queue[1:]
		for _, x := range neighbors(node[0], node[1], m, n, grid) {
			//fmt.Println("Neighbor:", x, "Value:", grid[x[0]][x[1]])
			_, ok := visited[x[0]+x[1]*m]
			if !ok {
				queue = append(queue, x)
			}
			visited[x[0]+x[1]*m] = true
		}
		//fmt.Println("Queue:", queue)
	}
	//fmt.Println("Counter:", counter)

	return counter
}

func maxAreaOfIsland(grid [][]int) int {
	// this avoids going multiple times over the same pixel
	//start := time.Now()

	var result int
	m, n := len(grid), len(grid[0])
	//fmt.Println(grid)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				// start search of neighbors
				count := breadthSearch(i, j, m, n, grid)
				if count > result {
					result = count
				}
			}
		}
	}

	//fmt.Println("Non-optimal algorithm took:", time.Since(start))

	return result
}

func main() {
	inputs := [][][]int{
		{
			{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		},
		{
			{0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	fmt.Println("Result:", maxAreaOfIsland(inputs[0]))
	fmt.Println("Result:", maxAreaOfIsland(inputs[1]))
}
