/*
Review Depth-first search
*/
package main

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func dfs(board [][]int, d, move []int) int {
	n := len(board)
	q := make([]int, 0)
	q = append(q, 1)
	d[1] = 0
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		for j := i + 1; j <= min(i+6, n*n); j++ {
			if move[j] > 0 {
				if d[move[j]] > d[i]+1 {
					q = append(q, move[j])
					d[move[j]] = d[i] + 1
				}
				continue
			}
			if d[j] > d[i]+1 {
				q = append(q, j)
				d[j] = d[i] + 1
			}

		}
	}

	if d[n*n] == n*n {
		return -1
	}
	return d[n*n]
}

func snakesAndLadders(board [][]int) int {
	var result int

	n := len(board)
	neigh := make([]int, n*n+1)
	move := make([]int, n*n+1)
	for i := 1; i <= n*n; i++ {
		neigh[i] = n * n
	}
	count, cur := 0, 1
	for i := n - 1; i >= 0; i-- {
		if count == 0 {
			for count < n {
				move[cur] = board[i][count]
				cur++
				count++
			}
		} else {
			for count > 0 {
				count--
				move[cur] = board[i][count]
				cur++
			}
		}
	}

	result = dfs(board, neigh, move)

	return result
}
