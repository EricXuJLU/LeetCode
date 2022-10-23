package main

import "sort"

// Day22 1235
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	dp := make([]int, len(startTime))
	type work struct {
		s int
		e int
		p int
	}
	works := make([]work, len(startTime))
	for i := 0; i < len(startTime); i++ {
		works[i].s = startTime[i]
		works[i].e = endTime[i]
		works[i].p = profit[i]
	}
	sort.Slice(works, func(i, j int) bool {
		return works[i].e < works[j].e
	})
	for i := range dp {
		dp[i] = works[i].p
	}
	ans := dp[0]
	for i := 1; i < len(works); i++ {
		//for j := 0; j < i; j++ {
		//	if works[j].e <= works[i].s && dp[i] < dp[j]+works[i].p {
		//		dp[i] = dp[j] + works[i].p
		//		continue
		//	}
		//}
		x := sort.Search(len(works), func(j int) bool {
			return works[j].e > works[i].s
		})
		if x > 0 {
			dp[i] = dp[x-1] + works[i].p
		}
		if dp[i] < dp[i-1] {
			dp[i] = dp[i-1]
		}
	}
	for _, v := range dp {
		if v > ans {
			ans = v
		}
	}
	return ans
}

// Day23 1768
func mergeAlternately(word1 string, word2 string) string {
	var ans []byte
	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			ans = append(ans, word1[i])
		}
		if i < len(word2) {
			ans = append(ans, word2[i])
		}
	}
	return string(ans)
}

// Day24 915
func partitionDisjoint(nums []int) int {
	leftMax, rightMin := make([]int, len(nums)), make([]int, len(nums))
	leftMax[0], rightMin[len(nums)-1] = nums[0], nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		leftMax[i] = max(leftMax[i-1], nums[i])
		rightMin[len(nums)-1-i] = min(rightMin[len(nums)-i], nums[len(nums)-1-i])
	}
	for i := 0; i < len(nums)-1; i++ {
		if leftMax[i] <= rightMin[i+1] {
			return i+1
		}
	}
	return len(nums)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}