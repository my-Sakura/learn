package main

func longestValidParentheses(s string) int {
	dp := make([]int, len(s))
	if len(s) < 2 {
		return 0
	}

	maxLen := 0
	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			dp[i] = 0
		} else {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		maxLen = max(maxLen, dp[i])
	}

	return maxLen
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
