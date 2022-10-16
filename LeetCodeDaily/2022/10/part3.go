package main

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

// Day 12 827
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, nums []int) int {
	// Solution 1
	//var grp []int
	//for p := head; p != nil; p = p.Next {
	//	grp = append(grp, p.Val)
	//}
	//used := make([]bool, len(grp))
	//for _, num := range nums {
	//	for i, val := range grp {
	//		if num == val {
	//			used[i] = true
	//		}
	//	}
	//}
	//ret := 0
	//for i := 1; i < len(used); i++ {
	//	if used[i-1] && !used[i] {
	//		ret++
	//	}
	//}
	//if used[len(used)-1] {
	//	ret++
	//}
	//return ret

	// Solution 2
	mp := make(map[int]bool)
	for _, v := range nums {
		mp[v] = true
	}
	cnt := 0
	for p, last := head, false; p != nil; p = p.Next {
		if mp[p.Val] && !last {
			cnt++
		}
		last = mp[p.Val]
	}
	return cnt
}

// Day 13 769
func maxChunksToSorted(arr []int) int {
	head, cnt := 0, 0
	for i := 0; i < len(arr); i++ {
		if day13(arr[head:i+1], head) {
			cnt++
			head = i + 1
		}
	}
	return cnt
}

func day13(arr []int, base int) bool {
	exist := make([]bool, len(arr))
	for _, a := range arr {
		if id := a - base; id >= len(arr) {
			return false
		} else {
			if exist[id] {
				return false
			} else {
				exist[id] = true
			}
		}
	}
	return true
}

// Day 15 1441
func buildArray(target []int, n int) []string {
	ps, pp := "Push", "Pop"
	cur := 1
	var ret []string
	for _, t := range target {
		for cur != t {
			ret = append(ret, ps, pp)
			cur++
		}
		ret = append(ret, ps)
		cur++
	}
	return ret
}
