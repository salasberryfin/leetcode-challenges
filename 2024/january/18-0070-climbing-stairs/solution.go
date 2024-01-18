package main

func climbStairs(n int) int {
	variations := make([]int, n+1)

	variations[0] = 1
	variations[1] = 1

	for i := 2; i <= n; i++ {
		variations[i] = variations[i-1] + variations[i-2]
	}

	return variations[n]
}
