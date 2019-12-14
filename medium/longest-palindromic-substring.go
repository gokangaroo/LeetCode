package main

import (
	"fmt"
	"math"
)

func main() {
	str := "eabcb"
	fmt.Println(longestPalindrome(str))
}

func longestPalindrome(s string) string {
	if s == "" {
		return s
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		len := int(math.Max(float64(len1), float64(len2)))
		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left int, right int) int {
	L, R := left, right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		// 向外拓展
		L--
		R++
	}
	return R - L - 1
}
