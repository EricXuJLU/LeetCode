package main

import (
	"fmt"
	"math"
	"math/big"
	"sort"
	"strconv"
	"strings"
)

// Day 1 1694
func reformatNumber(number string) string {
	sq := strings.Split(number, " ")
	number = strings.Join(sq, "")
	sq = strings.Split(number, "-")
	number = strings.Join(sq, "")
	sq = []string{}
	for i := 0; 3*i < len(number); i++ {
		sq = append(sq, number[3*i:func(a, b int) int {
			if a < b {
				return a
			}
			return b
		}(3*i+3, len(number))])
	}
	if len(sq[len(sq)-1]) == 1 {
		sq[len(sq)-1] = sq[len(sq)-2][2:] + sq[len(sq)-1]
		sq[len(sq)-2] = sq[len(sq)-2][:2]
	}
	number = strings.Join(sq, "-")
	return number
}

// Day 2 777
func canTransform(start string, end string) bool {
	var sl, sr, el, er []int
	for i, ch := range []byte(start) {
		if ch == 'L' {
			sl = append(sl, i)
		}
		if ch == 'R' {
			sr = append(sr, i)
		}
	}
	for i, ch := range []byte(end) {
		if ch == 'L' {
			el = append(el, i)
		}
		if ch == 'R' {
			er = append(er, i)
		}
	}
	if len(start) != len(end) || len(sl) != len(el) || len(sr) != len(er) {
		return false
	}
	for i := 0; i < len(sl); i++ {
		if sl[i] < el[i] {
			return false
		}
	}
	for i := 0; i < len(sr); i++ {
		if sr[i] > er[i] {
			return false
		}
	}
	// l r 次序不矛盾
	sq := strings.Split(start, "X")
	slr := strings.Join(sq, "")
	eq := strings.Split(end, "X")
	elr := strings.Join(eq, "")
	if strings.Compare(slr, elr) != 0 {
		return false
	}
	return true
}

// Day 3 1784
func checkOnesSegment(s string) bool {
	flag0 := false
	for i := 1; i < len(s); i++ {
		if s[i] == '1' && flag0 {
			return false
		}
		if s[i] == '0' {
			flag0 = true
		}
	}
	return true
}

// Day 4 921
func minAddToMakeValid(s string) int {
	ret, cnt := 0, 0
	for _, ch := range s {
		if ch == '(' {
			if cnt < 0 {
				ret += -cnt
				cnt = 1
			} else {
				cnt++
			}
		} else {
			cnt--
		}
	}
	ret += int(math.Abs(float64(cnt)))
	return ret
}

// Day 5 811
func subdomainVisits(cpdomains []string) []string {
	dMap := make(map[string]int)
	for _, s := range cpdomains {
		reps := strings.Split(s, " ")
		dms := strings.Split(reps[1], ".")
		cnt, _ := strconv.Atoi(reps[0])
		for i := 0; i < len(dms); i++ {
			key := strings.Join(dms[i:], ".")
			dMap[key] += cnt
			//fmt.Println("****")
		}
	}
	var ret []string
	for key, val := range dMap {
		ret = append(ret, fmt.Sprintf("%d %s", val, key))
	}
	return ret
}

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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Day 11 1790
func areAlmostEqual(s1 string, s2 string) bool {
	p1, p2 := -1, -1
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if p1 == -1 {
				p1 = i
			} else if p2 == -1 {
				p2 = i
			} else { // over 2
				return false
			}
		}
	}
	if p1 == -1 && p2 == -1 { // only 0
		return true
	}
	if p1 == -1 || p2 == -1 { // only 1
		return false
	}
	// actually 2
	if s1[p1] == s2[p2] && s1[p2] == s2[p1] {
		return true
	}
	return false
}

func main() {
	//fmt.Println(reformatNumber("1-23-45 6"))	//"123-456"
	//fmt.Println(reformatNumber("123 4-567"))	//"123-45-67"
	//fmt.Println(reformatNumber("123 4-5678"))	//"123-456-78"
	//fmt.Println(reformatNumber("12"))			//"12"

	//fmt.Println(canTransform("RXXLRXRXL", "XRLXXRRLX"))	//True
	//fmt.Println(canTransform("RL", "LR"))	//False
	//fmt.Println(checkOnesSegment("1001"))
	//fmt.Println(checkOnesSegment("110"))

	//fmt.Println(minAddToMakeValid("())"))	//1
	//fmt.Println(minAddToMakeValid("((("))	//3
	//fmt.Println(minAddToMakeValid("()))(("))	//4

	//fmt.Println(subdomainVisits([]string{"9001 discuss.leetcode.com"}))
	//fmt.Println(subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))

	//fmt.Println(maxAscendingSum([]int{10, 20, 30, 5, 10, 50}))
	//fmt.Println(maxAscendingSum([]int{10, 20, 30, 40, 50}))
	//fmt.Println(maxAscendingSum([]int{12, 17, 15, 13, 10, 11, 12}))
	//fmt.Println(maxAscendingSum([]int{100, 10, 1}))

	//fmt.Println(advantageCount([]int{2, 7, 11, 15}, []int{1, 10, 4, 11}))
	//fmt.Println(advantageCount([]int{12, 24, 8, 32}, []int{13, 25, 32, 11}))

	//fmt.Println(scoreOfParentheses("()"))
	//fmt.Println(scoreOfParentheses("(())"))
	//fmt.Println(scoreOfParentheses("()()"))
	//fmt.Println(scoreOfParentheses("(()(()))"))

	//fmt.Println(minSwap([]int{1, 3, 5, 4}, []int{1, 2, 3, 7}))
	//fmt.Println(minSwap([]int{0, 3, 5, 8, 9}, []int{2, 1, 4, 6, 9}))

	//fmt.Println(threeEqualParts([]int{1,0,1,0,1}))
	//fmt.Println(threeEqualParts([]int{1,1,0,1,1}))
	//fmt.Println(threeEqualParts([]int{1,1,0,0,1}))
	fmt.Println(threeEqualParts([]int{1, 1, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1}))
}
