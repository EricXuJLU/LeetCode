package main

import "sort"

// 1 6220
func averageValue(nums []int) int {
	sum, cnt := 0, 0
	for _, v := range nums {
		if v/6*6 == v {
			sum += v
			cnt++
		}
	}
	if cnt == 0 {
		return 0
	}
	return sum / cnt
}

// 2 6221
func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	mp := make(map[string]Creator)
	for i := 0; i < len(creators); i++ {
		c := mp[creators[i]]
		c.Name = creators[i]
		c.BFL += views[i]
		c.Ids = append(c.Ids, ids[i])
		c.Views = append(c.Views, views[i])
		mp[creators[i]] = c
	}
	var crs []Creator
	for _, v := range mp {
		sort.Sort(v)
		crs = append(crs, v)
	}
	sort.Slice(crs, func(i, j int) bool {
		return crs[i].BFL > crs[j].BFL
	})
	var ans [][]string
	maxBfl := crs[0].BFL
	for _, c := range crs {
		if c.BFL != maxBfl {
			break
		}
		ans = append(ans, []string{c.Name, c.Ids[0]})
	}
	return ans
}

type Creator struct {
	Name  string
	BFL   int
	Ids   []string
	Views []int
}

func (c Creator) Len() int {
	return len(c.Ids)
}

func (c Creator) Less(i, j int) bool {
	if c.Views[i] == c.Views[j] {
		return c.Ids[i] < c.Ids[j]
	}
	return c.Views[i] > c.Views[j]
}

func (c Creator) Swap(i, j int) {
	c.Ids[i], c.Views[i], c.Ids[j], c.Views[j] = c.Ids[j], c.Views[j], c.Ids[i], c.Views[i]
}

// 3 6222
func makeIntegerBeautiful(n int64, target int) int64 {
	x := int64(0)
	for bl := int64(1); ; bl *= 10 {
		if bitSum(n) <= target {
			return x
		}
		k := bl * 10
		del := k - n%k
		x += del
		n += del
	}
}

func bitSum(n int64) int {
	ans := 0
	for n != 0 {
		ans += int(n % 10)
		n /= 10
	}
	return ans
}

// 4 6223

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// PERSONAL VERSION
//func treeQueries(root *TreeNode, queries []int) []int {
//	mp, mp2 := make(map[int][]int), make(map[int][]int)
//	initMap(root, 0, &mp, &mp2)
//	var ans []int
//	for _, q := range queries {
//		m, ori := 0, mp[root.Val][1]
//		if mp[q][0]+mp[q][1] < ori {
//			ans = append(ans, ori)
//			continue
//		}
//		for _, v := range mp2[mp[q][0]] {
//			if v != q {
//				m = max(m, mp[v][0]+mp[v][1])
//			} else {
//				m = max(m, mp[v][0]-1)
//			}
//		}
//		ans = append(ans, m)
//	}
//	return ans
//}

func initMap(r *TreeNode, depth int, mp *map[int][]int, mp2 *map[int][]int) int {
	if r == nil {
		return -1
	}
	x := (*mp2)[depth]
	x = append(x, r.Val)
	(*mp2)[depth] = x
	tmp := []int{depth, 0}
	tmp[1] = max(initMap(r.Left, depth+1, mp, mp2), initMap(r.Right, depth+1, mp, mp2)) + 1
	(*mp)[r.Val] = tmp
	return tmp[1]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// OFFICIAL VERSION
func treeQueries(root *TreeNode, queries []int) []int {
	height := map[*TreeNode]int{} // 每棵子树的高度
	var getHeight func(*TreeNode) int
	getHeight = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		height[node] = 1 + max(getHeight(node.Left), getHeight(node.Right))
		return height[node]
	}
	getHeight(root)

	res := make([]int, len(height)+1) // 每个节点的答案
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, depth, restH int) {
		if node == nil {
			return
		}
		depth++
		res[node.Val] = restH
		dfs(node.Left, depth, max(restH, depth+height[node.Right]))
		dfs(node.Right, depth, max(restH, depth+height[node.Left]))
	}
	dfs(root, -1, 0)

	for i, q := range queries {
		queries[i] = res[q]
	}
	return queries
}
