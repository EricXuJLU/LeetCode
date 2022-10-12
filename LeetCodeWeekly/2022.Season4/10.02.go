package main

import (
	"math"
	"strings"
)

// 6192
func commonFactors(a int, b int) int {
	xiao := b
	if a < b {
		xiao = a
	}
	cnt := 0
	for i := 1; i <= xiao; i++ {
		if a/i*i == a && b/i*i == b {
			cnt++
		}
	}
	return cnt
}

// 6193
func maxSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	if n == 0 {
		return 0
	}
	maxOfAll := math.MinInt32
	for i := 1; i <= m-1-1; i++ {
		for j := 1; j <= n-1-1; j++ {
			maxOfThis := math.MinInt32
			lastShaLou := grid[i][j]
			//_, capability := lib.MinElement([]float32{float32(i), float32(m-1-i), float32(j), float32(n-1-j)})
			for k := 1; k <= 1; k++ {
				for l := j - k; l <= j+k; l++ {
					lastShaLou += grid[i-k][l] + grid[i+k][l]
				}
				if lastShaLou > maxOfThis {
					maxOfThis = lastShaLou
					//fmt.Println(maxOfThis, i, j, k)
				}
			}
			if maxOfThis > maxOfAll {
				maxOfAll = maxOfThis
			}
		}
	}
	return maxOfAll
}

// 6194 1 536870911
func minimizeXor(num1 int, num2 int) int {
	cnt1, cnt2 := 0, 0
	for num2 != 0 {
		cnt2 += num2 & 1
		num2 >>= 1
	}
	var bin []int
	for num1 != 0 {
		bin = append(bin, num1&1)
		cnt1 += num1 & 1
		num1 >>= 1
	}
	for len(bin) < 32 {
		bin = append(bin, 0)
	}
	ret := 0
	if cnt1 -= cnt2; cnt1 < 0 {
		// 末尾0转1
		for i := 0; i < len(bin) && cnt1 < 0; i++ {
			if bin[i] == 0 {
				bin[i] = 1
				cnt1++
			}
		}
	} else {
		// 末尾1转0
		for i := 0; i < len(bin) && cnt1 > 0; i++ {
			if bin[i] == 1 {
				bin[i] = 0
				cnt1--
			}
		}
	}
	for i := 0; i < len(bin); i++ {
		ret += bin[i] << i
	}
	return ret
}

// 6195
func deleteString(s string) int {
	dp := make([]int, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		index := deleteIndex(s[i:])
		if len(index) == 0 {
			dp[i] = 1
		} else {
			for _, j := range index {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[i+j]+1)))
			}
		}
	}
	return dp[0]
}

func deleteIndex(s string) []int {
	var ret []int
	for i := 1; 2*i <= len(s); i++ {
		if strings.Compare(s[:i], s[i:2*i]) == 0 {
			ret = append(ret, i)
		}
	}
	return ret
}
