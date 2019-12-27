package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	tmp := ""
	max := 0
	for i := 0; i < len(s); i++ {
		index := s[i : i+1]
		if !strings.Contains(tmp, index) { //** 这里也能通过空间换时间进行优化 **//
			tmp += index
		} else {
			if len(tmp) > max {
				max = len(tmp)
			}

			// 优化方案2: 用时0ms,人傻了
			//for i := 0; i < len(tmp); i++ {
			//	if tmp[i:i+1] != index {
			//		continue
			//	} else {
			//		tmp = tmp[i+1 : len(tmp)]
			//		break
			//	}
			//}

			// 优化方案1: 实际没有太大提升
			//for tmp[0:1] != index {
			//	// 去掉头
			//	tmp = tmp[1:len(tmp)]
			//}
			//tmp = tmp[1:len(tmp)]

			//**可优化
			for strings.Contains(tmp, index) {
				// 去掉头
				tmp = tmp[1:len(tmp)]
			}
			//**

			tmp += index
		}
	}
	if len(tmp) > max {
		max = len(tmp)
	}
	return max
}
