package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(calculate(" 3/2 "))
}

func calculate(s string) (ans int) {
	/*
	   官方题解
	       1. 总体思路, 转换成+号, append栈
	       1.1 负号直接负数存, 但是* / 需要对栈顶数字进行处理
	       2. 保存栈顶数字和当前数字间的符号
	*/
	// 去除两边的空格
	s = strings.TrimSpace(s)
	stack := make([]int, 0, len(s))
	mid := '+'
	num := 0
	for i, v := range s {
		if v == ' ' {
			continue
		}
		isDigital := v >= '0' && v <= '9'
		if isDigital {
			num = num*10 + int(v-'0')
		}
		if !isDigital || i == len(s)-1 {
			switch mid {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -1*num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			mid = v
			num = 0
		}
	}
	for _, v := range stack {
		ans += v
	}
	return
}
