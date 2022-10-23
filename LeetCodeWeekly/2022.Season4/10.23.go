package main

import (
	"sort"
	"strconv"
	"strings"
)

//1 6214
func haveConflict(event1 []string, event2 []string) bool {
	se1s := strings.Split(event1[0], ":")
	se1e := strings.Split(event1[1], ":")
	ie1s := trans(se1s)
	ie1e := trans(se1e)
	se2s := strings.Split(event2[0], ":")
	se2e := strings.Split(event2[1], ":")
	ie2s := trans(se2s)
	ie2e := trans(se2e)
	if (ie1s<=ie2s && ie1e>=ie2s) || (ie2s<=ie1s && ie2e>=ie1s) {
		return true
	}
	return false
}

func trans(s []string) float32 {
	var ans float32
	tmp, _ := strconv.Atoi(s[0])
	ans += float32(tmp)
	tmp, _ = strconv.Atoi(s[1])
	ans += 0.01 * float32(tmp)
	return ans
}

//2 6224
func subarrayGCD(nums []int, k int) int {
	var can [][]int
	for i, last, tail := 0, 0, 0; i < len(nums); i++ {
		if nums[i]/k*k == nums[i] {
			tail++
			if i == len(nums)-1 {
				if last < tail {
					can = append(can, nums[last:tail])
				}
				last = i+1
				tail = i+1
			}
		} else {
			if last < tail {
				can = append(can, nums[last:tail])
			}
			last = i+1
			tail = i+1
		}
	}
	ans := 0
	for _, c := range can {
		ans += count(c, k)
	}
	return ans
}

func count(nums []int, k int) int {
	for i := range nums {
		nums[i] /= k
	}
	//
	ans := 0
	for i := 0; i < len(nums); i++ {
		curGDC := nums[i]
		for j := i; j < len(nums); j++ {
			if curGDC = GDC(nums[j], curGDC); curGDC == 1 {
				ans++
			}
		}
	}
	return ans
}

func GDC(a, b int) int {
	if b != 0 {
		return GDC(b, a%b)
	} else {
		return a
	}
}

//3 6216
func minCost(nums []int, cost []int) int64 {
	sum := make([]int, 1e6+1)
	for i := 0; i < len(sum); i++ {
		for  {
			
		}
	}
	return 0
}

//4 6217
func makeSimilar(nums []int, target []int) int64 {
	nums, target = QuChong(nums, target)
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] %2 != nums[j]%2 {
			return nums[i]%2==1
		}
		return nums[i] < nums[j]
	})
	sort.Slice(target, func(i, j int) bool {
		if target[i] %2 != target[j]%2 {
			return target[i]%2==1
		}
		return target[i] < target[j]
	})
	cha := 0
	jocnt := 0
	for i := 0; i < len(nums); i++ {
		if x := target[i] - nums[i]; x > 0 {
			cha += x
			if x %2 == 1 {
				jocnt++
				cha--
			}
		}
	}
	return int64(cha/2+jocnt)
}

// QuChong 去重并排序
func QuChong(a, b []int) ([]int, []int) {
	ma, mb := make(map[int]int), make(map[int]int)
	for _, v := range a {
		ma[v]++
	}
	for _, v := range b {
		mb[v]++
	}
	for k := range ma {
		x := min(ma[k], mb[k])
		ma[k] -= x
		mb[k] -= x
	}
	var aa, bb []int
	for k, v := range ma {
		for i := 0; i < v; i++ {
			aa = append(aa, k)
		}
	}
	for k, v := range mb {
		for i := 0; i < v; i++ {
			bb = append(bb, k)
		}
	}
	sort.Ints(aa)
	sort.Ints(bb)
	return aa, bb
}

func min(a, b int) int {
	if a < b  {
		return a
	}
	return b
}