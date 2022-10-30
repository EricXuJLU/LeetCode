package main

import "fmt"

func main() {
	//10.02
	//fmt.Println(commonFactors(12, 6))
	//fmt.Println(commonFactors(25, 30))
	//grid := [][]int{
	//	{644460,685523,	769890,	443393,	285711},
	//	{9006,	254379,	969969,	601034,	689057},
	//	{440814,157646,	571484,	325351,	493670},
	//	{569808,429678,	388256,	869440,	914745},
	//	{589854,183767,	503184,	511695,	980066},
	//	{33725,	397553,	968569,	423475,	75219}}
	//fmt.Println(maxSum(grid))
	//fmt.Println(minimizeXor(1, 536870911))
	//fmt.Println(deleteString("abcabcdabc")) //2
	//fmt.Println(deleteString("aaabaab"))    //4
	//fmt.Println(deleteString("aaaaa"))      //5

	//10.09
	//fmt.Println(hardestWorker(10, [][]int{{0, 3}, {2, 5}, {0, 9}, {1, 15}}))
	//fmt.Println(hardestWorker(26, [][]int{{1, 1}, {3, 7}, {2, 12}, {7, 17}}))
	//fmt.Println(robotWithString("azz"))
	//fmt.Println(robotWithString("bac"))
	//fmt.Println(robotWithString("bdda"))
	//fmt.Println(robotWithString("vzhofnpo")) //"fnohopzv"
	//fmt.Println(numberOfPaths([][]int{
	//	{5, 2, 4}, {3, 0, 5}, {0, 7, 2},
	//}, 3))
	//fmt.Println(numberOfPaths([][]int{
	//	{0, 0},
	//}, 5))
	//fmt.Println(numberOfPaths([][]int{
	//	{7, 3, 4, 9}, {2, 3, 6, 2}, {2, 3, 7, 0},
	//}, 1))

	//10.23
	//fmt.Println(subarrayGCD([]int{9,3,1,2,6,3}, 3))
	//fmt.Println(subarrayGCD([]int{4}, 7))
	//fmt.Println(makeSimilar([]int{8,12,6}, []int{2,14,10}))
	//fmt.Println(makeSimilar([]int{1,2,5}, []int{4,1,3}))
	//fmt.Println(makeSimilar([]int{1,1,1,1,1}, []int{1,1,1,1,1}))
	//fmt.Println(makeSimilar([]int{758,334,402,1792,1112,1436,1534,1702,1538,1427,720,1424,114,1604,564,120,578},
	//						[]int{1670,216,1392,1828,1104,464,678,1134,644,1178,1150,1608,1799,1156,244,2,892}))
	n7 := &TreeNode{7, nil, nil}
	n6 := &TreeNode{6, nil, nil}
	n5 := &TreeNode{5, nil, n7}
	n4 := &TreeNode{4, n6, n5}
	n2 := &TreeNode{2, nil, nil}
	n3 := &TreeNode{3, n2, nil}
	n1 := &TreeNode{1, n3, n4}
	fmt.Println(treeQueries(n1, []int{4}))
}
