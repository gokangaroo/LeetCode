package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "110"
	fmt.Println(numDecodings(s))
}

//动态规划dp[i]
func numDecodings2(s string) int {
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

func numDecodings(s string) int {
	res := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '0' { // 1.s[i-1]必须是1和2,不然就return 0
			if i == 0 {
				return 0
			}
			if s[i-1] == '1' || s[i-1] == '2' {
				if i == 1 {
					res[i] = 1
					continue
				} else {
					res[i] = res[i-2]
					continue
				}
			} else {
				return 0
			}
		}
		// 2. 正常场景
		if i == 0 {
			res[0] = 1
		} else {
			now, _ := strconv.ParseInt(s[i-1:i+1], 0, 64)
			if now > 10 && now <= 26 { //需要考虑s[i-1]与s[i]是否可以一起编码
				if i == 1 {
					res[1] = 2
				} else {
					res[i] = res[i-1] + res[i-2]
				}
			} else {
				res[i] = res[i-1]
			}
		}
	}
	return res[len(s)-1]
}
