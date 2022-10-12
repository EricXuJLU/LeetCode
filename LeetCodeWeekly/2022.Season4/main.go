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
	fmt.Println(deleteString("abcabcdabc")) //2
	fmt.Println(deleteString("aaabaab"))    //4
	fmt.Println(deleteString("aaaaa"))      //5

	//10.09
	//fmt.Println(hardestWorker(10, [][]int{{0, 3}, {2, 5}, {0, 9}, {1, 15}}))
	//fmt.Println(hardestWorker(26, [][]int{{1, 1}, {3, 7}, {2, 12}, {7, 17}}))
	//fmt.Println(robotWithString("azz"))
	//fmt.Println(robotWithString("bac"))
	//fmt.Println(robotWithString("bdda"))
	//fmt.Println(robotWithString("vzhofnpo")) //"fnohopzv"
	fmt.Println(numberOfPaths([][]int{
		{5, 2, 4}, {3, 0, 5}, {0, 7, 2},
	}, 3))
	fmt.Println(numberOfPaths([][]int{
		{0, 0},
	}, 5))
	fmt.Println(numberOfPaths([][]int{
		{7, 3, 4, 9}, {2, 3, 6, 2}, {2, 3, 7, 0},
	}, 1))
}
