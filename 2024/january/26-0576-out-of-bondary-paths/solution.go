package main

const modulo int = 1000000007

func dfs(i, j, moves, rows, cols int, cache map[[3]int]int) int {
	if i < 0 || i == rows || j < 0 || j == cols {
		return 1
	}
	if moves == 0 {
		return 0
	}
	if _, ok := cache[[3]int{i, j, moves}]; ok {
		return cache[[3]int{i, j, moves}]
	}

	cache[[3]int{i, j, moves}] = ((dfs(i+1, j, moves-1, rows, cols, cache)+dfs(i-1, j, moves-1, rows, cols, cache))%modulo + (dfs(i, j+1, moves-1, rows, cols, cache)+dfs(i, j-1, moves-1, rows, cols, cache))%modulo) % modulo

	return cache[[3]int{i, j, moves}]
}

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	rows, cols := m, n
	cache := map[[3]int]int{}

	return dfs(startRow, startColumn, maxMove, rows, cols, cache)
}
