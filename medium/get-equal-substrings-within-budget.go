package main

import (
	"fmt"
)

func main() {
	fmt.Println(equalSubstring("krrgw", "zjxss", 19))
}

func equalSubstring(s string, t string, maxCost int) int {
	/*
	   1. 长度相同,同下标字符改变开销, ASCII差的绝对值
	   2. 最大开销maxCost, 大于则无法改变(消耗品)
	   3. 返回可以转换的最大长度

	   X 得到可用开销数组, 排序, 然后count++
	   Y 开销数组, 不排序, 而是求连续l最大值

	   FIX:
	   1. 相同的也算可以转换...
	   2. 要求的是连续长度...
	*/
	var costs = make([]int, len(s))
	for i := 0; i < len(s); i++ {
		costs[i] = transfer(s[i], t[i])
	}
	var start, end = 0, 0
	var maxLen = 0
	var sum = 0
	for end < len(s) {
		for sum+costs[end] > maxCost {
			// 减头保证buf够
			sum = sum - costs[start]
			start++
		}
		sum = sum + costs[end]
		if end-start+1 > maxLen {
			maxLen = end - start + 1
		}
		end++
	}
	return maxLen
}

func transfer(s, t byte) int {
	var a = int(s) - int(t)
	if a < 0 {
		a = -1 * a
	}
	return a
}
