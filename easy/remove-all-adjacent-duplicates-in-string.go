package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates("aaaaaaaaa"))
}

// 有相邻想等的"两个"字符就删除, 唯一注意的是, 删完需要注意填充的是否需要继续删
func removeDuplicates(S string) string {
	r := []byte(S)
	lenght := len(r)
	for i := 0; i < lenght-1; i++ { // 比较i和i+1
		for r[i] == r[i+1] && i < lenght-1 {
			if i+2 > lenght-1 { // 越界
				return string(r[:i])
			}
			copy(r[i:], r[i+2:])
			lenght = lenght - 2
			if i > 0 { //退一格继续
				i--
			}
		}
	}
	return string(r[:lenght])
}

// 官方题解,栈
func removeDuplicates2(S string) string {
	stack := make([]byte, 0, len(S))
	for i := range S {
		if len(stack) > 0 && stack[len(stack)-1] == S[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, S[i])
		}
	}
	return string(stack)
}
