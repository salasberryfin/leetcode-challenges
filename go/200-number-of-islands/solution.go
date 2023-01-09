/*
Keys:
	- Breadth first search (can be solved with Depth first search, also)
	- For each item, check each neighbor
	- The most important part is not going over same node more than once
		* keeping a `visited` slice is very time consuming on larger grids
		* the alternative is to, when one item has been processed, change it to a byte other than [0, 1]
		* now, we can discard those items just by value and don't have to check if contains(slice, value) which is O(n)
Notes:
	- The getIslandNeighbors() could be inside the bfs function. It is a standalone function
	just to make testing easier when I wasn't able to solve it.
*/
package main

var moves = [4][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func getIslandNeighbors(grid *[][]byte, r, c int) [][]int {
	var neighbors = [][]int{}
	for _, move := range moves {
		newC := c + move[0]
		newR := r + move[1]
		if 0 <= newR && newR < len(*grid) && 0 <= newC && newC < len((*grid)[0]) {
			if (*grid)[newR][newC] == '1' {
				neighbors = append(neighbors, []int{newR, newC})
				(*grid)[newR][newC] = '2'
			}
		}
	}

	return neighbors
}

func bfs(grid [][]byte, r, c int) {
	var q = [][]int{}
	q = append(q, []int{r, c})
	grid[r][c] = '2'
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		q = append(q, getIslandNeighbors(&grid, current[0], current[1])...)
	}
}

func numIslands(grid [][]byte) int {
	var islands int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				islands++
				bfs(grid, i, j)
			}
		}
	}

	return islands
}
