package main

import (
	"math/big"
	"sort"
)

// Day 6 927
func threeEqualParts(arr []int) []int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	if sum == 0 {
		return []int{0, 2}
	}
	if sum/3*3 != sum {
		return []int{-1, -1}
	}
	fst, scd, thd, cnt := -1, -1, -1, sum/3
	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 {
			cnt++
			if cnt >= sum/3 {
				if fst == -1 {
					fst = i
				} else if scd == -1 {
					scd = i
				} else {
					thd = i
				}
				cnt = 0
			}
		}
	}
	v3 := arr[thd:]
	l := len(arr) - thd
	v1 := arr[fst : fst+l]
	v2 := arr[scd : scd+l]
	if IsSame(v1, v2) && IsSame(v2, v3) {
		return []int{fst + l - 1, scd + l}
	}
	return []int{-1, -1}
}

func IsSame(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func caculate(arr []int) *big.Int {
	ret := big.NewInt(0)
	for i := len(arr) - 1; i >= 0; i-- {
		now := big.NewInt(int64(arr[i]))
		p := int64(len(arr) - 1 - i)
		pow := big.NewInt(p)
		pow = pow.Exp(big.NewInt(2), pow, big.NewInt(0))
		now = now.Mul(now, pow)
		ret = ret.Add(ret, now)
		//ret += arr[i] * 1<<(len(arr)-1-i)
	}
	return ret
}

// Day 7 1800
func maxAscendingSum(nums []int) int {
	dp := make([]int, len(nums))
	if len(nums) == 0 {
		return 0
	} else {
		for i := range nums {
			dp[i] = nums[i]
		}
	}
	ret := 0
	for i := 1; i < len(nums); i++ {
		j := i - 1
		if nums[j] < nums[i] {
			dp[i] = nums[i] + dp[j]
		}
	}
	for i := range nums {
		if dp[i] > ret {
			ret = dp[i]
		}
	}
	return ret
}

// Day 8 870
func advantageCount(nums1 []int, nums2 []int) []int {
	n2s := Ns{}
	for k, v := range nums2 {
		n2s = append(n2s, N{k: k, v: v})
	}
	sort.Sort(n2s)
	//fmt.Println(n2s)
	sort.Ints(nums1)
	ret := make([]int, len(nums1))
	flag := make([]bool, len(nums1))
	flag0 := make([]bool, len(nums1))
	for i1, i2 := 0, 0; i1 < len(nums1) && i2 < n2s.Len(); {
		if nums1[i1] > n2s[i2].v {
			ret[n2s[i2].k] = nums1[i1]
			flag[i1] = true
			flag0[n2s[i2].k] = true
			i1++
			i2++
			continue
		} else {
			i1++
		}
	}
	//fmt.Println(ret)
	//fmt.Println(flag)
	//fmt.Println(flag0)
	for i1, i0 := 0, 0; i1 < len(nums1) && i0 < n2s.Len(); {
		if !flag[i1] && !flag0[i0] {
			ret[i0] = nums1[i1]
			i1++
			i0++
			continue
		}
		if flag[i1] {
			i1++
		}
		if flag0[i0] {
			i0++
		}
	}
	return ret
}

type N struct {
	k int
	v int
}
type Ns []N

func (ns Ns) Less(i, j int) bool {
	return ns[i].v < ns[j].v
}

func (ns Ns) Swap(i, j int) {
	x := ns[i]
	ns[i] = ns[j]
	ns[j] = x
}

func (ns Ns) Len() int { return len(ns) }

// Day 9 856
func scoreOfParentheses(s string) int {
	//fmt.Println(s)
	if len(s) == 0 {
		return 0
	}
	if s == "()" {
		return 1
	}
	//if s[0] == '(' && s[len(s)-1] == ')' {
	//	return 2 * scoreOfParentheses(s[1:len(s)-1])
	//}
	sum, cnt, cut := 0, 0, 0
	for i, ch := range s {
		if ch == '(' {
			cnt++
		} else {
			cnt--
			if cnt == 0 {
				if i+1-cut > 2 && s[cut] == '(' && s[i] == ')' {
					sum += 2 * scoreOfParentheses(s[cut+1:i])
				} else {
					sum += scoreOfParentheses(s[cut : i+1])
				}
				cut = i + 1
			}
		}
	}
	return sum
}

// Day 10 801
func minSwap(nums1 []int, nums2 []int) int {
	dp := make([][2]int, len(nums1))
	dp[0][0], dp[0][1] = 0, 1
	for i := 1; i < len(nums1); i++ {
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] { //同位递增
			dp[i][0] = dp[i-1][0]
			if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] { //同位、错位都递增
				dp[i][0] = min(dp[i-1][1], dp[i][0])
				dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + 1
			} else { //仅同位递增
				dp[i][1] = dp[i-1][1] + 1
			}
		} else { //仅错位递增
			dp[i][0] = dp[i-1][1]
			dp[i][1] = dp[i-1][0] + 1
		}
	}
	return min(dp[len(nums1)-1][0], dp[len(nums1)-1][1])
}

