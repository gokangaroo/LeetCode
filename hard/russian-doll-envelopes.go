package main

import (
	"fmt"
	"sort"
)

func main() {
	envelopes := [][]int{
		{5, 40},
		{6, 8},
		{6, 70},
		{7, 9},
		{2, 30},
		{9, 11},
		{13, 12},
	}
	fmt.Println(maxEnvelopes(envelopes))
}

//动态规划思想  不过子问题的解决需要往前一次比较 双子序列
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] <= envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	res := make([]int, len(envelopes))
	max := 0
	for i := 0; i < len(envelopes); i++ {
		dpMax := 0
		for j := i; j >= 0; j-- {
			if envelopes[j][0] < envelopes[i][0] && envelopes[j][1] < envelopes[i][1] {
				if res[j] > dpMax {
					dpMax = res[j] // 这其实是个dp, =max(0-i)+1
				}
			}
		}

		res[i] = dpMax + 1 // res[i]表示自己作为最外层, 最多能包多少个.
		if res[i] > max {
			max = res[i]
		}
	}

	return max
}
