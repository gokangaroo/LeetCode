package main

import (
	"fmt"
)

func main() {
	s := "226"
	fmt.Println(numDecodings(s))
}

//动态规划dp[i]
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	var (
		pre  = 1
		curr = 1
	)
	for i := 1; i < len(s); i++ {
		tmp := curr
		//除非10，20合理 其余不可解码
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				curr = pre
			} else {
				return 0
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] >= '1' && s[i] <= '6') {
			curr = curr + pre
		}
		pre = tmp
	}
	return curr
}
