package main

import (
	"fmt"
	"math"
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
