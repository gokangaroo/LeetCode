package main

import (
	"fmt"
	"math"
)

func main() {
	num1 := 10
	num2 := 3
	fmt.Println(divide(num1, num2))
}

func divide(upnum int, downnum int) int {
	navigate := (upnum ^ downnum) < 0
	result := 0
	if downnum == 1 {
		return upnum
	}
	if upnum == downnum {
		return 1
	}
	if downnum == math.MinInt32 {
		return 0
	}
	if upnum == math.MinInt32 {
		if downnum == -1 {
			return math.MaxInt32
		} else {
			// go没有溢出
			upnum = If(downnum > 0, math.MaxInt32-downnum+1, math.MinInt32-downnum)
			result = 1
		}
	}
	upnum = If(upnum < 0, 0-upnum, upnum)
	downnum = If(downnum < 0, 0-downnum, downnum)
	var i uint32
	for i = 31; i != 4294967295; i-- {
		if (upnum >> i) >= downnum {
			result += (1 << i)
			upnum -= (downnum << i)
		}
	}
	return If(navigate, 0-result, result)

}

//模拟一个三元运算符
func If(condition bool, trueVal, falseVal int) int {
	if condition {
		return trueVal
	}
	return falseVal
}
