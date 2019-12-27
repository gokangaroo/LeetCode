package main

import (
	"fmt"
	"math"
)

func main() {
	str := "babad"
	fmt.Println(longestPalindrome(str))
}

func longestPalindrome(s string) string {
	if s == "" {
		return s
	}
	start, end, length := 0, 0, 0
	for i := 0; i < len(s)-1; i++ {
		len1 := expandAroundCenter(s, i, i)
		if s[i] == s[i+1] {
			len2 := expandAroundCenter(s, i, i+1)
			length = int(math.Max(float64(len1), float64(len2)))
		} else {
			length = len1
		}
		if length > end-start+1 { //新记录
			start = i - (length-1)/2
			end = i + length/2
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
