package main

// Day17 904
func totalFruit(fruits []int) int {
	dp := make([]int, len(fruits))
	dp[0] = 1
	lastdiff, ret := 0, 1
	for i := 1; i < len(fruits); i++ {
		if fruits[i] == fruits[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			if fruits[lastdiff] == fruits[i-1] {
				lastdiff = i - 1
				dp[i] = dp[i-1] + 1
			} else {
				if fruits[lastdiff] == fruits[i] {
					dp[i] = dp[i-1] + 1
					lastdiff = i - 1
				} else {
					dp[i] = i - lastdiff
					lastdiff = i - 1
				}
			}
		}
		if dp[i] > ret {
			ret = dp[i]
		}
	}
	return ret
}
