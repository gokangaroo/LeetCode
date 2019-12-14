package main

import (
	"fmt"
)

func main() {
	demo := []string{"aa", "a"}
	fmt.Println(longestCommonPrefix(demo))
}

func longestCommonPrefix(strs []string) string {
	result := ""
	length := len(strs)
	if length == 0 {
		return result
	}
	len := len(strs[0])
	// 这一段代码捕获代码异常
	additional(len, length, strs, &result)
	return result
}

func recoverPanic() {
	if r := recover(); r != nil {
	}
}

func additional(len int, length int, strs []string, result *string) {
	defer recoverPanic()
	for j := 0; j < len; j++ {
		for i := 0; i < length-1; i++ {
			if strs[i][j] == strs[i+1][j] {
			} else {
				panic("直接返回吧, 后面我不想看了")
			}
		}
		//说明这个字符j是公共前缀的一部分
		*result += string(strs[0][j])
	}
}
