package main

func dfs(grid [][]byte, row, col int) {
	moves := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	q := [][]int{}
	q = append(q, []int{row, col})
	grid[row][col] = '2'
	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		// iterate over each neighbor that is a '1'
		for _, move := range moves {
			newRow := current[0] + move[0]
			newCol := current[1] + move[1]
			// check if new position is inside the grid
			if newRow < len(grid) && newRow >= 0 && newCol < len(grid[0]) && newCol >= 0 && grid[newRow][newCol] == '1' {
				q = append(q, []int{newRow, newCol})
				grid[newRow][newCol] = '2' // update to a value different than 1 to avoid adding it again to the queue
			}
		}
	}
}

func numIslands(grid [][]byte) int {
	islands := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				islands++
				dfs(grid, i, j)
			}
		}
	}

	return islands
}
