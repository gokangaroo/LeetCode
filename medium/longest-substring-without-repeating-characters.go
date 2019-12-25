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
		if !strings.Contains(tmp, index) {
			tmp += index
		} else {
			if len(tmp) > max {
				max = len(tmp)
			}
			for strings.Contains(tmp, index) {
				// 去掉头
				tmp = tmp[1:len(tmp)]
			}
			tmp += index
		}
	}
	if len(tmp) > max {
		max = len(tmp)
	}
	return max
}
