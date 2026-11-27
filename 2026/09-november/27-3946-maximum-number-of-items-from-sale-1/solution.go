package main

func maximumSaleItems(items [][]int, budget int) int {
	n := len(items)

	frees := make([]int, n)
	for i := range n {
		for j := range n {
			if i != j && items[j][0]%items[i][0] == 0 {
				frees[i]++
			}
		}
	}

	dp := make([]int, budget+1)
	for w := 1; w <= budget; w++ {
		dp[w] = -1
	}

	for i := range n {
		cost := items[i][1]
		gain := 1 + frees[i]
		for w := budget; w >= cost; w-- {
			if dp[w-cost] != -1 {
				if dp[w-cost]+gain > dp[w] {
					dp[w] = dp[w-cost] + gain
				}
			}
		}
	}

	minPrice := items[0][1]
	for i := 1; i < n; i++ {
		if items[i][1] < minPrice {
			minPrice = items[i][1]
		}
	}

	ans := 0
	for w := 0; w <= budget; w++ {
		if dp[w] != -1 {
			extraCopies := (budget - w) / minPrice
			if dp[w]+extraCopies > ans {
				ans = dp[w] + extraCopies
			}
		}
	}

	return ans
}
