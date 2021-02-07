package main

import "fmt"

func main() {
	fmt.Println(maxScore([]int{1, 2, 3, 4, 5, 6, 1}, 3))
}

func maxScore(cardPoints []int, k int) int {
	/*
	   一行卡牌, 整数点数
	   每次从首尾选择一张牌, 拿k张 (1 <= k <= cardPoints.length)
	   k张点数和, 得到最大点数.

	   1. 如果用递归dfs很简单, 但是性能不行, 10^5数组太长, 会爆栈, 会超时
	   2. 逆向思维, 计算len(cardPoints)-k个连续数字的和最小值..
	   3. 或者将前k个数拼到后k个数的后面(如果够的话), 然后计算k个连续数字的最大值
	*/
	var sum = 0
	for i := 0; i < len(cardPoints)-k; i++ {
		sum = sum + cardPoints[i]
	}
	var windows = sum
	var min = windows
	for i := len(cardPoints) - k; i < len(cardPoints); i++ {
		sum = sum + cardPoints[i]
		windows = windows + cardPoints[i] - cardPoints[i+k-len(cardPoints)]
		if windows < min {
			min = windows
		}
	}
	return sum - min
}
