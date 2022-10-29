package main

// Day26 862
func shortestSubarray(nums []int, k int) int {
	return -1
}

// Day27 1822
func arraySign(nums []int) int {
	ans := 1
	for _, v := range nums {
		if v == 0 {
			return 0
		}
		if v < 0 {
			ans *= -1
		}
	}
	return ans
}

// Day28 907
// 排序，按最小元素分治递归
func sumSubarrayMins(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0] % int(1e9+7)
	}
	id := getMinIndex(arr)
	return ((sumSubarrayMins(arr[:id])+(id+1)*(len(arr)-id)*arr[id])%int(1e9+7) + sumSubarrayMins(arr[id+1:])) % int(1e9+7)
}

func getMinIndex(arr []int) int {
	id := 0
	for i, v := range arr {
		if v < arr[id] {
			id = i
		}
	}
	return id
}

// Day29 1773
func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	key := 0
	switch ruleKey {
	case "type":
		key = 0
	case "color":
		key = 1
	case "name":
		key = 2
	}
	mp := make(map[string]int)
	for _, it := range items {
		mp[it[key]]++
	}
	return mp[ruleValue]
}

// Day30 784
func letterCasePermutation(s string) []string {
	var ans []string
	cnt := 0
	for _, ch := range s {
		if isLetter(ch) {
			cnt++
		}
	}
	psb := 1 << cnt
	for gan := 0; gan < psb; gan++ {
		var w []byte
		for b, i := 0, 0; i < len(s); i++ {
			if isLetter(int32(s[i])) {
				bit := 1 << b
				bit &= gan
				w = append(w, shift(s[i], bit != 0))
				b++
			} else {
				w = append(w, byte(s[i]))
			}
		}
		ans = append(ans, string(w))
	}
	return ans
}

func shift(ch uint8, hc bool) byte {
	if ch < 'a' {
		ch -= 'A'
	} else {
		ch -= 'a'
	}
	if hc {
		return ch + 'A'
	} else {
		return ch + 'a'
	}
}

func isLetter(ch int32) bool {
	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
		return true
	}
	return false
}
