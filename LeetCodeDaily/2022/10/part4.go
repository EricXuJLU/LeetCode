package main

import (
	"math"
	"strconv"
)

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

// Day18 902
func atMostNGivenDigitSet(digits []string, n int) int {
	var dig []int
	for _, s := range digits {
		x, _ := strconv.Atoi(s)
		dig = append(dig, x)
	}
	var dn []int
	for x := n; x != 0; x /= 10 {
		dn = append(dn, x%10)
	}
	var ans = -1
	for bl := 1; bl <= n; bl *= 10 {
		ans += PowInt(len(dig), int(math.Log10(float64(bl))))
		if n/(10*bl) == 0 {
			pdn := len(dn) - 1
			for i := 0; i < len(dig) && pdn >= 0; i++ {
				if dn[pdn] > dig[i] {
					if i == len(dig)-1 {
						ans += PowInt(len(dig), pdn+1)
						return ans
					}
				} else if dn[pdn] == dig[i] {
					ans += i * PowInt(len(dig), pdn)
					pdn--
					if pdn < 0 {
						ans++
						return ans
					}
					i = -1
				} else {
					ans += i * PowInt(len(dig), pdn)
					return ans
				}
			}
			break
		}
	}
	return ans
}

func PowInt(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Day19 1700
func countStudents(students []int, sandwiches []int) int {
	eated := make([]bool, len(students))
	var xs, smz = 0, 0
	for smz < len(sandwiches) {
		for lazy := xs; ; {
			if eated[xs] == true {
				xs = (xs + 1) % len(students)
				if xs == lazy {
					return len(sandwiches) - smz
				}
				continue
			}
			if students[xs] == sandwiches[smz] {
				eated[xs] = true
				xs = (xs + 1) % len(students)
				smz++
				break
			} else {
				xs = (xs + 1) % len(students)
				if xs == lazy {
					return len(sandwiches) - smz
				}
			}
		}
	}
	return 0
}
