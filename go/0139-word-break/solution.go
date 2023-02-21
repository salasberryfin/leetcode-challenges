package main

func wordBreak(s string, wordDict []string) bool {
	dp := []bool{}
	for i := 0; i <= len(s); i++ {
		dp = append(dp, false)
	}
	dp[len(s)] = true
	for i := len(s) - 1; i >= 0; i-- {
		for _, w := range wordDict {
			currentLength := i + len(w)
			if currentLength <= len(s) && s[i:currentLength] == w {
				dp[i] = dp[currentLength]
			}
			if dp[i] {
				break
			}
		}
	}

	return dp[0]
}
