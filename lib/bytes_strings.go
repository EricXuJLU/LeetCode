package lib

import "strings"

// IsDigital 判断字符是否为0~9的数值
func IsDigital(bt byte) bool {
	return bt-'0'>=0 && bt-'0'<10
}

// IsInteger 判断字符串是否为一个整数（不允许带空格，允许+00000）
func IsInteger(str string) bool {
	bts := []byte(str)
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

// IsNumber1 判断s是否为一个可以由科学计数法表示的数（小数点前后可以单空，如".1", "1."）
func IsNumber1(s string) bool {
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

// IsNumber2 判断s是否为一个可以由科学计数法表示的数（小数点前后皆不能空，如"1.0"）
func isNumber2(s string) bool {
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