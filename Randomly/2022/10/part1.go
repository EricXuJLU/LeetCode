package main

// 2789
func missingNumber(nums []int) int {
	n := len(nums) + 1
	should := (0 + n - 1) * n / 2
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return should - sum
}

// 187
func findRepeatedDnaSequences(s string) []string {
	mp := make(map[string]int)
	for i := 0; i <= len(s)-10; i++ {
		mp[s[i:i+10]]++
	}
	var ret []string
	for k, v := range mp {
		if v > 1 {
			ret = append(ret, k)
		}
	}
	return ret
}
