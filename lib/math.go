package lib

import "math"

// CommonFactors 两个正整数的所有公因数
func CommonFactors(a int, b int) []int {
	xiao := b
	if a < b {
		xiao = a
	}
	var ret []int
	for i := 1; i <= xiao; i++ {
		if a/i*i == a && b/i*i == b {
			ret = append(ret, i)
		}
	}
	return ret
}

// MaxElement 求最大元素
func MaxElement(a []float32) (index int, ans float32) {
	if len(a) == 0 {
		return -1, 0
	}
	index = 0
	for i := 1; i < len(a); i++ {
		if a[i] > a[index] {
			index = i
		}
	}
	return index, a[index]
}

// MinElement 求最小元素
func MinElement(a []float32) (index int, ans float32) {
	if len(a) == 0 {
		return -1, 0
	}
	index = 0
	for i := 1; i < len(a); i++ {
		if a[i] < a[index] {
			index = i
		}
	}
	return index, a[index]
}

// PowInt int型的幂次函数，注意返回值也在int范围
func PowInt(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
