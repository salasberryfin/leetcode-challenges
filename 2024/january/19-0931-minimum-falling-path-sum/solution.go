package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dfs(row, col int, matrix [][]int, cache map[[2]int]int) int {
	if row == len(matrix) {
		return 0
	}
	if val, ok := cache[[2]int{row, col}]; ok {
		return val
	}
	if col < 0 || col >= len(matrix) {
		return 1000000
	}
	firstMin := min(dfs(row+1, col-1, matrix, cache), dfs(row+1, col, matrix, cache))
	result := matrix[row][col] + min(firstMin, dfs(row+1, col+1, matrix, cache))
	cache[[2]int{row, col}] = result

	return result
}

func minFallingPathSum(matrix [][]int) int {
	result := 1000000

	for i := 0; i < len(matrix); i++ {
		result = min(result, dfs(0, i, matrix, make(map[[2]int]int)))
	}

	return result
}
