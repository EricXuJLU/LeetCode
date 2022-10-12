package main

import (
	"sort"
)

// 1 2432
func hardestWorker(n int, logs [][]int) int {
	l := Table(logs)
	sort.Sort(l)
	for i := len(l) - 1; i > 0; i-- {
		l[i][1] -= l[i-1][1]
	}
	sort.Sort(l)
	return l[len(l)-1][0]
}

type Table [][]int

func (t Table) Len() int {
	return len(t)
}
func (t Table) Less(i, j int) bool {
	if t[i][1] != t[j][1] {
		return t[i][1] < t[j][1]
	} else {
		return t[i][0] > t[j][0]
	}
}
func (t Table) Swap(i, j int) {
	x := t[i]
	t[i] = t[j]
	t[j] = x
}

// 2 6201
func findArray(pref []int) []int {
	arr := make([]int, len(pref))
	arr[0] = pref[0]
	sum := pref[0]
	for i := 1; i < len(arr); i++ {
		arr[i] = pref[i] ^ sum
		sum = arr[i] ^ sum
	}
	return arr
}

// 3 6202
func robotWithString(s string) string {
	t := make([]byte, 0, len(s))
	cnt := make([]int, 26)
	for _, c := range s {
		cnt[c-'a']++
	}
	var ret []byte
	for i, p := 0, 0; i < len(s); i++ {
		t = append(t, s[i])
		cnt[s[i]-'a']--
		p = next(&cnt, p)
		if p == -1 {
			break
		}
		for len(t) > 0 && int(t[len(t)-1]) <= p+'a' {
			ret = append(ret, t[len(t)-1])
			t = t[:len(t)-1]
		}
	}
	for len(t) != 0 {
		ret = append(ret, t[len(t)-1])
		t = t[:len(t)-1]
	}
	return string(ret)
}

func next(v *[]int, p int) int {
	if (*v)[p] != 0 {
		return p
	}
	for i := p; i < len(*v); i++ {
		if (*v)[i] != 0 {
			return i
		}
	}
	return -1
}

// 4 6203
func numberOfPaths(grid [][]int, k int) int {
	mod := int(1e9 + 7)
	m, n := len(grid), len(grid[0])
	dp := make([][][]int, m) // dp[m][n][k]
	for i := 0; i < m; i++ {
		t1 := make([][]int, n)
		for j := 0; j < n; j++ {
			t2 := make([]int, k)
			t1[j] = t2
		}
		dp[i] = t1
	}
	dp[0][0][grid[0][0]%k] = 1
	for i := 1; i < m; i++ {
		for j := 0; j < k; j++ {
			dp[i][0][j] = dp[i-1][0][(j-grid[i][0]+100*k)%k] % mod
		}
	}
	for i := 1; i < n; i++ {
		for j := 0; j < k; j++ {
			dp[0][i][j] = dp[0][i-1][(j-grid[0][i]+100*k)%k] % mod
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			for l := 0; l < k; l++ {
				pre := (l - grid[i][j] + 100*k) % k
				dp[i][j][l] = (dp[i][j-1][pre] + dp[i-1][j][pre]) % mod
			}
		}
	}
	return dp[m-1][n-1][0]
}
