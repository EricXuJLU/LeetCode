package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isNumber(" -1e-16 "))
	//fmt.Println(IsInteger("+"))
}

// IsDigital 判断字符是否为0~9的数值
func IsDigital(ch byte) bool {
	return ch-'0'>=0 && ch-'0'<10
}

// IsInteger 判断字符串是否为一个整数（不允许带空格，允许+00000）
func IsInteger(string2 string) bool {
	bts := []byte(string2)
	if len(bts) == 0 {
		return false
	}
	if bts[0] == '+' || bts[0] == '-' {
		bts = bts[1:]
		if len(bts) == 0 {
			return false
		}
	}
	for _, bt := range bts {
		if !IsDigital(bt) {
			return false
		}
	}
	return true
}

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	bt := []byte(s)
	if len(bt) == 0 {
		return false
	}
	if bt[0] == '+' || bt[0] == '-' {
		bt = bt[1:]
	}
	for i, ch := range bt {
		if ch == 'e' || ch == 'E' {
			if !IsInteger(string(bt[i+1:])) {
				return false
			}
			bt = bt[:i]
			break
		}
	}
	if len(bt) == 0 {
		return false
	}
	for i, ch := range bt {
		if ch == '.' {
			bt1 := bt[:i]
			bt2 := bt[i+1:len(bt)]
			if len(bt1) == 0 && len(bt2) == 0 {
				return false
			}
			for _, c := range bt1 {
				if !IsDigital(c) {
					return false
				}
			}
			for _, c := range bt2 {
				if !IsDigital(c) {
					return false
				}
			}
			return true
		}
		if !IsDigital(ch) {
			return false
		}
	}
	return true
}