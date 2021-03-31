package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}

func subsetsWithDup(nums []int) (ans [][]int) {
	/*
	   官方题解一: 迭代,TODO
	   官方题解二: 递归, 递归好理解, 主要看递归
	*/
	sort.Ints(nums)
	var t []int
	// 是否选择前面的下标对应的数字, 当前数字的下标
	var dfs func(bool, int)
	dfs = func(choosePre bool, cur int) {
		// 终止条件
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), t...))
			return
		}
		// 递归条件
		// 1. 当前位置不选,递归cur+1
		dfs(false, cur+1)
		// 如果前者没选, 现者与前者同, 则跳过(剪枝,不需要考虑选择的情况)
		if !choosePre && cur > 0 && nums[cur-1] == nums[cur] {
			return
		}
		// 2. 选择当下, 递归cur+1
		t = append(t, nums[cur])
		dfs(true, cur+1)
		// 回溯的时候, 需要把t进行去尾部
		t = t[:len(t)-1]
	}
	// 从第一个数字开始!
	dfs(false, 0)
	return
}
